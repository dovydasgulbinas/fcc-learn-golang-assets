package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("Pausing writing to a channel for 5000 ms ...")
		time.Sleep(time.Millisecond * 5000)
		messages <- "ping"
		fmt.Println("I wrote to channel 'ping'")
	}()

	fmt.Println("I wait for channel to finish writting a message")

	msg := <-messages // code will get blocked from execution until there is any information in the channel
	fmt.Printf("I Received a message '%v'\n", msg)

	fmt.Println("Program has exited")
}
