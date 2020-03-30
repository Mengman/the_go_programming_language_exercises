package main

// Exercise 5.14: Use the breadthFirst function to explore a different structure. For example,
// you could use the course dependencies from the topoSort example (a directed graph), the file
// system hierarchy on your computer (a tree), or a list of bus or subway routes downloaded from
// your city government’s website (an undirected graph).

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	breadthFirst(travelDirectory, os.Args[1:])
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

func travelDirectory(root string) []string {
	fmt.Println(root)

	list, err := doTravel(root)
	if err != nil {
		log.Print(err)
	}

	return list
}

func doTravel(root string) (folders []string, err error) {
	items, err := ioutil.ReadDir(root)
	if err != nil {
		return
	}

	for _, item := range items {
		if item.IsDir() {
			folders = append(folders, filepath.Join(root, item.Name()))
		}
	}

	return
}
