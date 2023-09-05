package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("(goroutine) I will do nothing and I will not update any channels.")
	}()

	fmt.Println("I will wait for channel updates...")
	fmt.Println(
		"Deadlock will occur, because all goroutines have finished execution and messages have not been updated.",
	)
	msg := <-messages                              // code will get blocked from execution until there is any information in the channel
	fmt.Printf("I Received a message '%v'\n", msg) // it will not print this
	fmt.Println("Program has exited")              // it will not exit
}
