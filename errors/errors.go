package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	time    time.Time
	message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.time, e.message)
}

func run() error {
	return &MyError{
		time.Now(),
		"Something went wrong!!!",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("42.1")

	if err != nil {
		fmt.Printf("Error Message: %s\n", err)
		return
	}
	fmt.Println("Converted int", i)
}
