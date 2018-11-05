package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println("the sum of all the multiples of 3 or 5 below 1000: ", sum)

	forSum := 1
	for forSum < 1000 {
		forSum += forSum
	}
	fmt.Println(forSum)

	fmt.Println("Square root of 1024 is ", sqrt(1024))

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
