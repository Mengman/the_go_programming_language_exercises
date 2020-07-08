package main

import (
	"fmt"
	"time"
	"flag"
)

func buildPipeline(stageNum int64) (input, out chan int) {
	out = make(chan int)
	start := out
	for i := int64(0); i < stageNum; i++ {
		in := out
		out = make(chan int)
		go func(in, out chan int) {
			for i := range in {
				out <- i
			}
			close(out)
		}(in, out)
	}
	return start, out
}

func main() {
	n := flag.Int64("stage", 1000, "pipeline stage number")
	flag.Parse()
	
	in, out := buildPipeline(*n)
	tick := time.Now()
	
	in <- 1
	<-out
	
	fmt.Printf("transit value through %d stages pipeline takes %v\n", *n, time.Since(tick))
}
