package main

import (
	"fmt"
	"time"
)

func main() {
	// Channels ----------------------------------------------
	fmt.Println("Channels Section")

	messages := make(chan string)
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	// Channel Buffering -------------------------------------
	fmt.Println("Channel Buffering Section")

	// make channels of string buffering up to 2 values.
	messages = make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Channel Synchronization -------------------------------
	fmt.Println("Channel Synchronization Section")

	done := make(chan bool, 1)
	go worker(done)

	<-done

}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Print("done")

	done <- true
}
