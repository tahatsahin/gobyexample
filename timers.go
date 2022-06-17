package main

import (
	"fmt"
	"time"
)

func main() {
	// wait 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// this channel blocks on the timer's channel C until it sends a
	// value indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// wait can be performed by time.Sleep as well. one reason a timer
	// may be useful is that you can cancel the timer before it fires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
