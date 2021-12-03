package ch01

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func printSum(v1, v2 *value) {
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()

	time.Sleep(1 * time.Second)
	v2.mu.Lock()
	defer v2.mu.Unlock()

	fmt.Println("sum=%v\n", v1.value+v2.value)
}
