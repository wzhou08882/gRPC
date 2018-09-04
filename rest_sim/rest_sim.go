// ### Globals and imports
package main

import (
	"fmt"
	"log"
	"os"
	// "os/exec"
	"time"
	"strings"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pub"
	"github.com/go-mangos/mangos/protocol/sub"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

// newPublisherSocket creates a new pub socket from the passed-in URL, 
// and starts listening on this socket.
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

func subscribe(socket mangos.Socket, topic string) error {
        err := socket.SetOption(mangos.OptionSubscribe, []byte(topic))
        if err == nil {
                err = socket.SetOption(mangos.OptionRecvDeadline, 3600*time.Second)
        }
        return err
}

func receive(socket mangos.Socket) (string, error) {
        message, err := socket.Recv()
        return string(message), err
}

// topics that the server will use for sending messages.
func runServer(url string, topics []string) {
	// Create the publisher socket.
	socket, err := newPublisherSocket(url)
	if err != nil {
		log.Fatalf("Cannot listen on %s: %s\n", url, err.Error())
	}
        // loop using following commands as message
	var cmds = [][]string{{"ls", "-l"},
                              {"ifconfig", "-a", "-v"},
                              {"netstat", "-v", "-s"},
                              {"netstat", "-v", "-l"},
                              {"cat", "/proc/cpuinfo"},
                              {"cat", "/proc/meminfo"},
                              {"ps", "all"},
                              {"pstree", "-l"}}
	i := 0;
	var cmd_line string
	for true {
		cl := len(cmds);
		for _, topic := range topics {
			time.Sleep(5 * time.Second)
			cmd_line = strings.Join(cmds[i]," ")
			fmt.Printf("CMD <%s> ==> CPE <%s>\n", cmd_line, topic)
			err = publish(socket, topic, cmd_line);
			i++;
			i = i % cl;
			if err != nil {
				log.Fatalf("Cannot publish message for topic %s: %s\n", topic, err.Error())
			}
		}
	}
}

// Putting it all together...
func main() {

	// The socket URL for publish.
	url_pub := "tcp://localhost:56565"
	// The socket URL for subscribe.
	url_sub := "tcp://localhost:56575"

	go func() {
		fmt.Println("Start Waitting CPEs Response:")
		socket, err := newSubscriberSocket(url_sub);
		if err != nil {
			log.Fatalf("Can't wait at:%s: %s\n", url_sub, err.Error())
		}
		for _, topic := range os.Args[1:] {
			err = subscribe(socket, topic)
			if err != nil {
				log.Fatalf("Cannot get CPE response: %s\n", err.Error())
			}
		}
		tl := len(os.Args[1])
                var p int
		for true {
		        message, err := receive(socket);
		        if err != nil {
                                log.Fatalf("Error message: %s\n", err.Error())
			}
			p = tl
                        cpe := message[:p]
                        p++
                        msg := message[p:]
			fmt.Println("======== RSP MSG Start =========")
			fmt.Printf("From:  %s\n", cpe)
			fmt.Printf("%s", msg)
			fmt.Println("============ MSG END ===========")
			fmt.Println("")
                }
	}()
	// Without parameters, the process starts will exit
	l := len(os.Args);
        if l == 1 {
		fmt.Println("Usage: ./rest_sim [<cpe uuid>]");
		return;
	} else {
		fmt.Println("Starting REST Simulator server")
		runServer(url_pub, os.Args[1:])
                fmt.Println("Server ends.")
	}
}

