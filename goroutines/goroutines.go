package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// wg.Done() will execute only after the completion of loop
func say(s string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		duration := 100 * time.Microsecond
		time.Sleep(duration)
		fmt.Println(s)
		fmt.Println("Function running in background, Duration:", duration)
	}
}

// the last method only executes after the execution of
// first two methods as we have added two methods to wait group
func main() {
	wg.Add(2)
	go say("Hello")
	go say("World")
	wg.Wait()

	wg.Add(1)
	say("Another world")
}
