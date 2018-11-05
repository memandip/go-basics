package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(100, 10, 20))
	fmt.Println(sqrt(100))
}

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
