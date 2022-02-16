package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// You will able to see that all worker pools are running concurrently
	// This way we are making use of the available cores in our machine
	// If you run this we will able to see the CPU usage almost 100% on all core

	// This is just the gist of how worker pools works
	// It does not guarantee that fibonacci will be calculated in parallel or accurate
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
