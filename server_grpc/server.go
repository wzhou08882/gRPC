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
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"github.com/go-mangos/mangos"
        "github.com/go-mangos/mangos/protocol/sub"
	"github.com/go-mangos/mangos/protocol/pub"
        "github.com/go-mangos/mangos/transport/ipc"
        "github.com/go-mangos/mangos/transport/tcp"
	"google.golang.org/grpc/testdata"
	pb "cpedemo/cpe_demo"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	port       = flag.Int("port", 10000, "The server port")
)

type cpeDemoServer struct {
	registeredCpe map[string]string 
	msgCount int32;
	mu sync.Mutex; 
	reportSocket mangos.Socket 
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
                // A second socket option avoids that clients wait forever when they receive no messages.
                err = socket.SetOption(mangos.OptionRecvDeadline, 150*time.Second)
        }
        return err
}

func receive(socket mangos.Socket) (string, error) {
        message, err := socket.Recv()
        return string(message), err
}

func (s *cpeDemoServer) Register(ctx context.Context, cpeinfo *pb.CpeInfo) (*pb.MsgStruc, error) {
	fmt.Println("gRPC Server: Register() requested!")
	s.mu.Lock()
	s.msgCount = 1
	//s.registeredCpe[cpeinfo.GetCpeName()] = cpeinfo.GetCpeUuid()
	var rsp pb.MsgStruc
	s.msgCount++
	rsp.MsgId = s.msgCount
	rsp.Message = "CPE Registered"
	s.mu.Unlock()
	return &rsp, nil
}

func (s *cpeDemoServer) Report(ctx context.Context, cpe2cm *pb.Cpe2Cm) (*pb.MsgStruc, error) {
	fmt.Println("gRPC Server: Report() requested!")
	uuid := cpe2cm.GetCpeUuid()
	rsp := cpe2cm.GetResInfo()
	// fmt.Printf("get CPE<%s>\t RSP<%s>", uuid, rsp)
	fmt.Printf("Pub Response to REST, MSG#:%d, CPE<%s>\n", rsp.GetMsgId(), uuid)
	err := publish(s.reportSocket, uuid, rsp.GetMessage());
	if err != nil {
		log.Fatalf("Cannot publish RPS to CPE<%s>:%s\n", uuid, err.Error())
		return nil, err
	}
	var msg pb.MsgStruc
        msg.MsgId = rsp.MsgId
        msg.Message = "Got It!"
	return &msg, nil
}

func (s *cpeDemoServer) Subscribe(cpeinfo *pb.CpeInfo, stream pb.CpeDemo_SubscribeServer) error {
        fmt.Println("## gRPC Server: Subscribe() requested!")
	fmt.Println("## I will keep sending config commands when got from REST API ##")
	// The socket URLs.
        url_sub := "tcp://localhost:56565"

        // create a subscriber socket.
	socket, err := newSubscriberSocket(url_sub)
	if err != nil {
                log.Fatalf("Cannot dial into %s: %s\n", url_sub, err.Error())
        }
	topic := cpeinfo.GetCpeUuid()
	err = subConfigInfo(socket, topic)
	tl := len(topic)
	var p int
        for true {
                message, err := receive(socket)
                if err != nil {
                        log.Fatalf("Error receiving message: %s\n", err.Error())
			continue;
                }
		p = tl
		cpe := message[:p]
                p++
		msg := message[p:]
		s.mu.Lock()
		s.msgCount++
		fmt.Printf("gRPC SERVER: No.%d MSG<%s> to CPE<%s>\n", s.msgCount,
			msg, cpe)
                err = stream.Send(&pb.MsgStruc{
                        MsgId:  s.msgCount,
                        Message:  msg});
		s.mu.Unlock()
                if err != nil {
                    return err
                }

        }
	return nil

}

func newServer() *cpeDemoServer {
	s := &cpeDemoServer{}
	// The publish socket URLs.
        url_pub := "tcp://localhost:56575"

        // create publish socket for response Report
        sktPub, er := newPublisherSocket(url_pub)
        if er != nil {
                log.Fatalf("Cannot listen on %s: %s\n", url_pub, er.Error())
        }
        s.reportSocket = sktPub
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCpeDemoServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
