package main

import (
	"fmt"
	"math"
)

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func abs(f myFloat) float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := myFloat(math.Sqrt(10))
	fmt.Println(f.abs())
}
