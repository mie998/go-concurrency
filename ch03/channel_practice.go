package ch03

import (
	"bytes"
	"fmt"
	"os"
)

func deadlock() {
	stringCh := make(chan string)
	go func() {
		for {
		}
		stringCh <- "hello"
	}()
}

func channelWithBuffer() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intCh := make(chan int, 4)
	go func() {
		defer close(intCh)
		defer fmt.Println("Producer done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "sending...: %d\n", i)
			intCh <- i
		}
	}()

	for integer := range intCh {
		fmt.Fprintf(&stdoutBuff, "Recieved: %d\n", integer)
	}
}

func channelLifetime() {
	chanOwner := func() <-chan int {
		resultCh := make(chan int, 5)
		go func() {
			defer close(resultCh)
			for i := 0; i < 5; i++ {
				resultCh <- i
			}
		}()
		return resultCh
	}

	resultCh := chanOwner()
	for result := range resultCh {
		fmt.Println("Received:", result)
	}
	fmt.Println("Done receiving!")
}
