package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"unicode"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	var contents []string
	contents = outline(contents, doc)
	for _, v := range contents {
		fmt.Printf("%s\n", v)
	}
}


func outline(contents []string, n *html.Node) []string {
	if n.Type == html.TextNode && !isWhiteSpace(n.Data){
		contents = append(contents, n.Data)
	}

	// when found script tag stop traverse it's child node
	if n.Data == "script" {
		return contents
	}

	for c:= n.FirstChild; c != nil; c = c.NextSibling {
		contents = outline(contents, c)
	}
	return contents
}

func isWhiteSpace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
