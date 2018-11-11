package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type ShapeInfo struct {
	area, perimter float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func getShapeInfo(s Shape) ShapeInfo {
	return ShapeInfo{s.area(), s.perimeter()}
}

func main() {
	circle := Circle{10}
	rect := Rectangle{10, 20}
	circleInfo := getShapeInfo(circle)
	rectInfo := getShapeInfo(rect)
	fmt.Println("Area of circle ", circleInfo.area)
	fmt.Println("Perimeter of circle ", circleInfo.perimter)

	fmt.Println("Area of rectange ", rectInfo.area)
	fmt.Println("Perimeter of rectange ", rectInfo.perimter)

	// with interface implementation
	var c Shape = Circle{10}
	var r Shape = Rectangle{10, 20}
	fmt.Printf("Thea area and perimeter of circle is %v and %v respectively.\n", c.area(), c.perimeter())
	fmt.Printf("Thea area and perimeter of rectangle is %v and %v respectively.\n", r.area(), r.perimeter())
}
