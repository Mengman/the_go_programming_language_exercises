package main

// Exercise 5.13: Modify crawl to make local copies of the pages it finds, creating directories as
// necessary. Don’t make copies of pages that come from a different domain. For example, if the
// original page comes from golang.org, save all files from there, but exclude ones from vimeo.com.

import (
	"fmt"
	"log"
	"os"

	"github.com/Mengman/the_go_programming_language_exercises/ch5/ex5.13/web"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := web.Extract(url, true)
	if err != nil {
		log.Print(err)
	}
	return list
}
