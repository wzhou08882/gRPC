/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"flag"
	"io"
	"log"
	"time"
        "fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
        "github.com/go-mangos/mangos"
        "github.com/go-mangos/mangos/protocol/sub"
        "github.com/go-mangos/mangos/protocol/pub"
        "github.com/go-mangos/mangos/transport/ipc"
        "github.com/go-mangos/mangos/transport/tcp"
	"github.com/google/uuid"
	pb "cpedemo/cpe_demo"
)

var (
	tls = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func printCpeInfo(cpeInfo *pb.CpeInfo) {
	fmt.Println("CPE Info:")
	fmt.Printf("\tCPE uuid: %s\n", cpeInfo.GetCpeUuid());
	fmt.Printf("\tCPE name: %s\n", cpeInfo.GetCpeName());
	fmt.Printf("\tGlobal IP: %s", cpeInfo.GetExternalIp());
	fmt.Printf("\tLocal IP: %s\n", cpeInfo.GetLocalIp());
        fmt.Printf("\tCPE other: %s\n", cpeInfo.GetOther());
}

func runRegister(client pb.CpeDemoClient, in *pb.CpeInfo) {
	log.Printf("gRPC Client: runRegister(): CPE<%s>\n", in.GetCpeUuid()) 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rsp, err := client.Register(ctx, in)
	if err != nil {
		log.Fatalf("%v.Register(_) = _, %v", client, err)
	}
	log.Printf("gRPC Client: Register result: %s", rsp.GetMessage())
}


func runSubscribe(client pb.CpeDemoClient, cpeinfo *pb.CpeInfo) {
	log.Printf("gRPC Client: runSubscrib(): CPE<%s>\n", cpeinfo.GetCpeUuid())
	log.Println("gRPC Client: Waiting for Config Commands <====")
        ctx, cancel := context.WithTimeout(context.Background(), 3600*time.Second)
        defer cancel()
        stream, err := client.Subscribe(ctx, cpeinfo)
        if err != nil {
                log.Fatalf("%v.Subscribe(_) = _, %v", client, err)
        }
	//log.Printf("gRPC Client: Subscribe result:", stream.GetMessage())
        var cpe2cm pb.Cpe2Cm
	cpe2cm.ResInfo = new(pb.MsgStruc)
	cpe2cm.CpeUuid = cpeinfo.GetCpeUuid()
	var strArray [] string
        for {
            in, err := stream.Recv()
            if err == io.EOF {
		continue
            }
            if err != nil {
                log.Fatalf("Failed to receive Config Info : %v", err)
            }
	    log.Println("gRPC Client Got:")
	    log.Printf("    MSG ID: %d", in.MsgId)
	    log.Printf("    Message: %s", in.Message)
            strArray = strings.Fields(in.Message)
	    rsp := executeCmd(strArray)
	    cpe2cm.ResInfo.MsgId = in.MsgId
	    cpe2cm.ResInfo.Message = rsp
	    msg, er := client.Report(ctx, &cpe2cm)
            if er != nil {
                log.Fatalf("Failed to Report result: %v", er)
	    }
	    log.Printf("client.Report() result: %s", msg.GetMessage())
        }
        stream.CloseSend()
}

func genUuid() string {
        id := uuid.New()
        return id.String()
}

func GetOutboundIP() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP.String()
}

func GetExternalIP() string {
        //cmd := "dig +short myip.opendns.com @resolver1.opendns.com"
        out, err := exec.Command("dig", "+short", "myip.opendns.com", 
			"@resolver1.opendns.com").Output()
        if err != nil {
                log.Fatalf("can't get IP:", err.Error())
                return ("Error!")
        }
        return string(out[:])
}

func createCpeInfo(cpeName string) *pb.CpeInfo {
	var cpeinfo pb.CpeInfo

	cpeinfo.CpeName = cpeName
	cpeinfo.CpeUuid = genUuid()
	cpeinfo.LocalIp = GetOutboundIP()
	cpeinfo.ExternalIp = GetExternalIP()
	cpeinfo.Other = ""
	return &cpeinfo
}

func executeCmd(cmd_line [] string) string {
	cmd := cmd_line[0]
	args := cmd_line[1:]
        out, err := exec.Command(cmd, args...).Output()
        if err != nil {
                log.Fatalf("Error on execution of : %s", cmd)
                return ("Error!")
        }
        return string(out[:])
}

func newPublisherSocket(url string) (mangos.Socket, error) {
        socket, err := pub.NewSocket()
        if err != nil {
                return nil, err
        }
        // Allow the use of either TCP or IPC.
        socket.AddTransport(ipc.NewTransport())
        socket.AddTransport(tcp.NewTransport())

        // Start listening.
        err = socket.Listen(url)
        if err != nil {
                return nil, err
        }

        return socket, nil
}

func publish(socket mangos.Socket, topic, message string) error {
        err := socket.Send([]byte(fmt.Sprintf("%s|%s", topic, message)))
        return err
}

func newSubscriberSocket(url string) (mangos.Socket, error) {
        socket, err := sub.NewSocket()
        if err != nil {
                return nil, err
        }
        socket.AddTransport(ipc.NewTransport())
        socket.AddTransport(tcp.NewTransport())
        err = socket.Dial(url)
        if err != nil {
                return nil, err
        }
        return socket, nil
}

func subConfigInfo(socket mangos.Socket, topic string) error {
        err := socket.SetOption(mangos.OptionSubscribe, []byte(topic))
        if err == nil {
                err = socket.SetOption(mangos.OptionRecvDeadline, 150*time.Second)
        }
        return err
}

func receive(socket mangos.Socket) (string, error) {
        message, err := socket.Recv()
        return string(message), err
}

func main() {
	if len(os.Args) == 1 {
		log.Printf("Usage: ./client <CPE NAME>")
		return
	}
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCpeDemoClient(conn)

	// Create CPE info
	var cpeInfo *pb.CpeInfo
        cpeInfo = createCpeInfo(os.Args[1]);
	printCpeInfo(cpeInfo);

        // run CPE Register
	runRegister(client, cpeInfo);

	// run CPE Subscribe
	runSubscribe(client, cpeInfo);
}
