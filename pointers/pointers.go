package main

import (
	"fmt"
)

func main() {
	var i, j = 10, 18
	p := &i
	fmt.Println(*p)
	*p = 20
	fmt.Println(i)

	p = &j
	fmt.Println(*p)
	*p = *p / 10
	fmt.Println(j)
}
