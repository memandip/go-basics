package main

import (
	"fmt"
	"time"
)

func main() {
	// make a buffered channel
	c := make(chan string, 2)

	// It will throw error because the channel is already full
	// it has the capacity of 2
	// c <- "Hello"
	// c <- "World"
	// c <- "World"

	go count("Hello", c)
	go count("world", c)
	go count("mandip", c)

	for {
		// when receiving from channels we receive two values
		// one is the value and the other is if the channel is still open
		msg, open := <-c

		if !open {
			break
		}
		fmt.Println(msg)
	}

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan<- string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	// channel should only be closed on the sender side
	// receiver side should not close the channel
	close(c)
}
