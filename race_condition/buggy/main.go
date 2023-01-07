package main

import (
	"fmt"
	"sync"
)

var x int = 0

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i += 1 {
		wg.Add(1)

		go worker(&wg)
	}

	// Wait until every worker calls wg.Done().
	wg.Wait()

	fmt.Println("x = %i", x)
}

func worker(wg *sync.WaitGroup) {
	for i := 0; i < 10_000; i += 1 {
		x += 1
	}

	wg.Done()
}
