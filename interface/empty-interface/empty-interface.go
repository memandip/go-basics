package main

import "fmt"

func main() {
	// empty interface
	var i interface{}
	describe(i)

	i = []string{"Mandip", "Tharu"}
	describe(i)

	i = "hello"
	i, ok := i.(string)
	describe(i)
	fmt.Println("OK", ok)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
