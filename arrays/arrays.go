package main

import (
	"fmt"
	"strings"
)

func main() {
	// var letters = []string{"a", "b", "c", "d", "e", "f"}
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	// lettersArray := [3]string{"a", "b", "c"}
	// fmt.Println(lettersArray)
	const limit int = 26
	a := make([]string, limit)

	fmt.Println(letterRunes)

	for i := 0; i < limit; i++ {
		a[i] = string(letterRunes[i])
	}
	fmt.Println(a)
	fmt.Println(strings.Join(a, ""))

	a = append(a, "A")

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)
	// var slice = primes[0:5]
	// fmt.Println(slice)

	// dynamicArray := make([]int, 10)
	// fmt.Println("Dynamic array", dynamicArray)

	var pow = []int{1, 2, 4, 8, 64, 128, 312, 624}
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}

}
