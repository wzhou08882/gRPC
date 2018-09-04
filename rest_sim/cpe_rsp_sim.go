// ### Globals and imports
package main

import (
	"fmt"
	"log"
	// "os"
	// "os/exec"
	"time"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pub"
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

// topics that the server will use for sending messages.
func runServer(url string, topic string) {
	// Create the publisher socket.
	socket, err := newPublisherSocket(url)
	if err != nil {
		log.Fatalf("Cannot listen on %s: %s\n", url, err.Error())
	}
        // loop using following commands as message
	var cmds = []string{"ls", "ifconfig", "route", "hostname", 
	                        "whoami", "df", "pstree", "free", 
				"date", "uptime"}
	i := 0;
	for true {
		cl := len(cmds);
		time.Sleep(3 * time.Second)
		fmt.Printf("CMD <%s>\t ==> CPE <%s>\n", cmds[i], topic)
		err = publish(socket, topic, cmds[i]);
		i++
		i = i % cl;
		if err != nil {
			log.Fatalf("Cannot publish message for topic %s: %s\n", topic, err.Error())
		}
	}
}

// Putting it all together...
func main() {

	// The socket URL for publish.
	url := "tcp://localhost:56575"

	fmt.Println("Starting CPE Response Simulator")
	runServer(url, "CpeResponse")
}

