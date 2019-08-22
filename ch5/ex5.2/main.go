package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	nodeCounter := make(map[string]int32)
	outline(nodeCounter, doc)
	for k, v := range nodeCounter {
		fmt.Printf("%s: %d\n", k, v)
	}
}


func outline(counter map[string]int32, n *html.Node) {
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}

	for c:= n.FirstChild; c != nil; c = c.NextSibling {
		outline(counter, c)
	}
}
