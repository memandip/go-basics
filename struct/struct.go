package main

import (
	"fmt"
	"strconv"
)

type Vertex struct {
	X int
	Y int
}

type Person struct {
	firstName, lastName, gender string
	age                         int
}

// value receiver
func (p Person) greet() string {
	return "Hi, I am " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// pointer receiver
func (p *Person) increaseAge() {
	p.age++
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{Y: 5}
	p1 = &v2
)

func main() {
	vertex := Vertex{10, 20}
	vertex.X = vertex.X * 5
	vertex.Y = vertex.Y * 12
	fmt.Println(vertex.X, vertex.Y)

	p := &vertex
	p.X = 1e10
	fmt.Println(vertex)
	// p1.X = 100
	fmt.Println(v1, v2, *p1)

	person := Person{"Mandip", "Tharu", "Male", 23}

	fmt.Println(person.greet())

	person.increaseAge()
	fmt.Println(person.greet())
}
