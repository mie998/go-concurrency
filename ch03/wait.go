package ch03

import (
	"fmt"
	"runtime"
	"sync"
)

func ex1() {
	var wg sync.WaitGroup
	salutation := "Hello,"
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(salutation)
	}()
	salutation = "hannya???"
	wg.Wait()
}

func ex2() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"Hello", "hannya???", "unchi"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func ex3() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"Hello", "hannya???", "unchi"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func ex4() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}

func main() {
	ex4()
}
