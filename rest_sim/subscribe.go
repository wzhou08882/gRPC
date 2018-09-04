package main

import (
	"fmt"
	"log"
	"os"
	"time"

	// For this example, we need the PUBSUB protocol as well as the ipc and tcp transports.
	// Unlike the PAIR protocol, PUBSUB actually consists of two protocols, PUB and SUB.
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/sub"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

// newSubscriberSocket creates a new sub socket from the passed-in URL, and dials
// into this socket.
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
		// A second socket option avoids that clients wait forever when they receive no messages.
		err = socket.SetOption(mangos.OptionRecvDeadline, 10*time.Second)
	}
	return err
}

// Receiving is nothing more than calling the socket's Recv() method. The magic happens
// through the socket option "OptionSubscribe" we set earlier. This option makes the
// socket ignore any message that does not start with the desired topic(s).
func receive(socket mangos.Socket) (string, error) {
	message, err := socket.Recv()
	return string(message), err
}

// Client setup is also easy.
func runClient(url string, topic string) {
	// First, we create a subscriber socket.
	socket, err := newSubscriberSocket(url)
	if err != nil {
		log.Fatalf("Cannot dial into %s: %s\n", url, err.Error())
	}
	err = subscribe(socket, topic)
	if err != nil {
		log.Fatalf("Cannot subscribe to topic %s: %s\n", topic, err.Error())
	}
	// Finally, we listen for new message and print out any that matches
	// one of the topics we subscribed to.
	for true {
		message, err := receive(socket)
		if err != nil {
			log.Fatalf("Error receiving message: %s\n", err.Error())
		}
		tl := len(topic)
		name := message[:tl]
		tl++
		msg := message[tl:]
		fmt.Printf("Client <%s> received <%s>\n", name, msg)
	}
}

// Putting it all together...
func main() {

	// The socket URL.
	url := "tcp://localhost:56565"


	// One or more parameters means this process is a client.
	fmt.Println(os.Args[1], "is starting.")
	runClient(url, os.Args[1])
	fmt.Println("Client", os.Args[1], "ends.")
}

