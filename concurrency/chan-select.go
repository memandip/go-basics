package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	// This way even if the channel 1 (c1) is ready to send data
	// it will have to wait two seconds before sending data
	// Because the channel 2 (c2) is not ready to send data and channels are blocking
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	// It will select the based on the channel which is ready to send data
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
