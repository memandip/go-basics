package main

import (
	"fmt"
)

type Shape struct {
	X, Y float64
}

// pointer receiver as method of Shape struct
func (s *Shape) scale(scale float64) Shape {
	s.X = s.X * scale
	s.Y = s.Y * scale
	return *s
}

// pointer receiver function
func scale(s *Shape, scale float64) Shape {
	s.X = s.X * scale
	s.Y = s.Y * scale
	return *s
}

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

	shape := Shape{2, 3}
	fmt.Println("Current shape values,", shape)
	fmt.Println("Scaled shape values,", shape.scale(2))
	fmt.Println("Scaled shape values,", scale(&shape, 2))
}
