package main

import (
	"fmt"
)

type Planet struct {
	name        string
	age         float64 //age in billions
	isHabitable bool
}

func (p Planet) String() string {
	return fmt.Sprintf("%v (%v billon years) (isHabitable = %v)\n", p.name, p.age, p.isHabitable)
}

// When we use fmt to print something it looks for the String method
// We can override the default behaviour my defining it in our package
func main() {
	a := Planet{"Earth", 4.5, true}
	b := Planet{"Mars", 5.2, false}
	fmt.Println(a, b)
}
