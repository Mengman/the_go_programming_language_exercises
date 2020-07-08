package main

import (
	"fmt"
	"os"
	"time"
)

func bounce(ping, pong chan int) {
	counter := 0
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case v := <-ping:
			counter++
			pong <- v
		case <-tick:
			fmt.Printf("channel throughput: %d/s\n", counter)
			counter = 0
		}
	}
}

func main() {

	ping := make(chan int)
	pong := make(chan int)

	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("abort test")
		abort <- struct{}{}
	}()

	go bounce(ping, pong)
	go bounce(pong, ping)

	ping <- 1

	<-abort
}
