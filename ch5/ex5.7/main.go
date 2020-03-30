package main

// Exercise 5.7: Develop startElement and endElement into a general HTML pretty-printer.
// Print comment nodes, text nodes, and the attributes of each element (<a href='...'>). Use
// short forms like <img/> instead of <img></img> when an element has no children. Write a
// test to ensure that the output can be parsed successfully. (See Chapter 11.)


import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"unicode"
)

func main() {
	for _, url := range os.Args[1:] {
		doc, err := loadHtmlPage(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "HTML pretty-printer: %v\n", err)
			continue
		}
		forEachNode(doc, startElement, endElement)
	}
}

func loadHtmlPage(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return nil, err
	}
	return doc, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c!= nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type != html.ElementNode {
		return
	}
	data := n.Data
	if isWhiteSpace(data) {
		data = ""
	}
	fmt.Printf("%*s<%s", depth*2, "", data)
	for _, a := range n.Attr {
		fmt.Printf(" %s='%v' ", a.Key, a.Val)
	}
	if n.FirstChild != nil {
		fmt.Printf(">\n")
	} else{
		fmt.Printf("/>\n")
	}
	depth++
}


func endElement(n *html.Node) {
	if n.Type != html.ElementNode {
		return
	}

	depth--
	if n.FirstChild != nil {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func isWhiteSpace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}