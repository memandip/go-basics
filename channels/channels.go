package main

import (
	"fmt"
)

func chanTest(c chan int, val int) {
	c <- 5 * val
}

// Channels are a typed conduit through which you can send and receive
// values with the channel operator, <-.
// channels are blocking by default
func main() {
	c := make(chan int)
	go chanTest(c, 10)
	go chanTest(c, 15)
	x, y := <-c, <-c
	fmt.Printf("X = %v, Y = %v\n", x, y)
}
