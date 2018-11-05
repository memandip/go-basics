package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Some other OS")
	}

	fmt.Println("When is Wednesday?")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("After 2 days")
	default:
		fmt.Println("Still many days")
	}
	fmt.Println("Time : ", time.Now().Hour())
}
