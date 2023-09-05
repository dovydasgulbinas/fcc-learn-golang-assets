package main

import "fmt"

func main() {
	myString := "hello"
	fmt.Printf("myString = %v\n", myString)

	myStringPtr := &myString
	fmt.Printf("myStringPtr = %v\n", myStringPtr)

	*myStringPtr = "world I was re-pointed"
	fmt.Printf("Changing value using pointers ...\n")
	fmt.Printf("myString = %v\n", myString)
}
