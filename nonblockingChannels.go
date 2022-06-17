package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receive, if there is a value available on channel
	// messages, then select will take the <-messages case with value
	// if not, it will immediately take the default case.
	select {
	case msg := <-messages:
		fmt.Println("received message ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	// non blocking send, msg cannot be sent to messages channel,
	// because the channel has no buffer and there is no receiver.
	// therefore default case is selected
	select {
	case messages <- msg:
		fmt.Println("sent message ", msg)
	default:
		fmt.Println("no message sent")
	}

	// multi cases before default, multi-way non-blocking select.
	// non blocking receives on both messages and signals.
	select {
	case msg := <-messages:
		fmt.Println("received message ", msg)
	case sig := <-signals:
		fmt.Println("received signal ", sig)
	default:
		fmt.Println("no activity")
	}
}
