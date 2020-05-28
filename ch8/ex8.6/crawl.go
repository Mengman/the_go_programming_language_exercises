package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/Mengman/the_go_programming_language_exercises/ch5/ex5.13/web"
)

var tokens = make(chan struct{}, 20)
var maxDepth int
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

func main() {
	flag.IntVar(&maxDepth, "depth", 3, "max crawl depth")
	flag.Parse()

	wg := new(sync.WaitGroup)

	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 0, wg)
	}

	wg.Wait()
}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth > maxDepth {
		return
	}

	tokens <- struct{}{}
	list, err := web.Extract(url, false, "")
	<-tokens
	if err != nil {
		log.Print(err)
	}

	for _, link := range list {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}
