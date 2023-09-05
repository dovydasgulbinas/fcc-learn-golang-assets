package main

import "fmt"

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color // the derefencing of c argument is done automatically
}

func main() {
	c := car{
		color: "white",
	}
	fmt.Println(c.color) // prints "white"
	c.setColor("black")
	fmt.Println(c.color) // prints "black"
}
