package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 3; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
		wg.Wait()
	}
}
