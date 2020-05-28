package main

import (
	"flag"
	"fmt"
	"github.com/Mengman/the_go_programming_language_exercises/ch5/ex5.13/web"
	"log"
	"strings"
)

var tokens = make(chan struct{}, 20)

func main() {
	var domain string
	flag.StringVar(&domain, "domain", "", "domain name")
	flag.Parse()

	if len(domain) < 1 {
		flag.Usage()
		return
	}

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- []string{domain} }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !strings.Contains(link, domain) {
				continue
			}

			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := web.Extract(url, true, "/media/tx-deepocean/Data/tmp")
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
