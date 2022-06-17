package main

import (
	"fmt"
	"time"
)

func main() {
	// Channels ----------------------------------------------
	fmt.Println("\n\nChannels Section")

	messages := make(chan string)
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	// Channel Buffering -------------------------------------
	fmt.Println("\n\nChannel Buffering Section")

	// make channels of string buffering up to 2 values.
	messages = make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Channel Synchronization -------------------------------
	fmt.Println("\n\nChannel Synchronization Section")

	done := make(chan bool, 1)
	go worker(done)

	<-done

	// Channel Directions ------------------------------------
	fmt.Println("\n\n\nChannel Directions Section")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Print("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
