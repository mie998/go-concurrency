package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var cadence *sync.Cond
var left, right int32
var peopleInHallway sync.WaitGroup

func takeStep() {
	cadence.L.Lock()
	defer cadence.L.Unlock()
	cadence.Wait()
}

func tryDir(dirName string, dir *int32, out *bytes.Buffer) bool {
	fmt.Fprintf(out, " %v", dirName)
	atomic.AddInt32(dir, 1)
	takeStep()
	if atomic.LoadInt32(dir) == 1 {
		fmt.Fprint(out, ". Success!")
		return true
	}
	takeStep()
	atomic.AddInt32(dir, -1)
	return false
}

func tryLeft(out *bytes.Buffer) bool {
	return tryDir("left", &left, out)
}

func tryRight(out *bytes.Buffer) bool {
	return tryDir("right", &right, out)
}

func walk(walking *sync.WaitGroup, name string) {
	var out bytes.Buffer
	defer func() { fmt.Println(out.String()) }()
	defer walking.Done()
	fmt.Fprintf(&out, "%v is styring to scoot.", name)

	for i := 0; i < 5; i++ {
		if tryLeft(&out) || tryRight(&out) {
			return
		}
	}
	fmt.Fprintf(&out, "\nI'm lost!\n")
}

func main() {
	cadence = sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
