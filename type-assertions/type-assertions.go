package main

import "fmt"

func main() {
	fmt.Println("Type Assertions")
	fmt.Println("===============")

	var i interface{} = []string{"Hello World!!!"}

	a, ok := i.([]string)
	fmt.Println(a)

	b, ok := i.(string)
	fmt.Printf("Value: %v, OK: %v\n", b, ok)

	c, ok := i.(float64)
	fmt.Println(c, ok)

	// it throws an error
	// d := i.(float64)
	// fmt.Println(d)
}
