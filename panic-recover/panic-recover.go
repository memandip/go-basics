package main

import (
	"fmt"
)

func getPanic() interface{} {
	panic("PANIC ERROR")
}

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Panic Recovered,", r)
	}
}

func main() {
	defer cleanup()
	getPanic()
}
