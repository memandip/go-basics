package main

import "fmt"

func main() {
	// defer fmt.Println("World")
	// fmt.Println("hello")

	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
