package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var a int
	var b string
	var c int64
	var d bool
	fmt.Printf("%v '%v' %v %v\n", a, b, c, d)
	fmt.Println("Hello World!!!")
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(randomString(10))
	fmt.Println(math.Pi)
	sum := add(10, 20)
	fmt.Println(sum)
	fmt.Println(split(sum))
	fmt.Println(swap("Hello", "World"))
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < 1; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
