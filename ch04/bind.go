package ch04

import (
	"fmt"
)

func chanOwner() <-chan int {
	results := make(chan int, 5)
	go func() {
		defer close(results)
		for i := 0; i <= 5; i++ {
			results <- i
		}
	}()
	return results
}

func consumer(results <-chan int) {
	for result := range results {
		fmt.Println("Received:", result)
	}
	fmt.Println("Done receiving!")
}
