package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var found = false

func main() {
	url := os.Args[1]
	id := os.Args[2]
	doc, err := loadHtmlPage(url)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ElementByID: %v\n", err)
		os.Exit(1)
	}
	node := ElementByID(doc, id)
	if node != nil {
		fmt.Printf("%s\n", node.Data)
	} else {
		fmt.Printf("Element with ID '%s' not found.\n", id)
	}
}

func loadHtmlPage(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return nil, err
	}
	return doc, nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, &id, findElementByID, nil)
}

func forEachNode(n *html.Node, id *string, pre, post func(n *html.Node, id *string) bool) *html.Node {
	if pre != nil {
		found = pre(n, id)
	}

	if !found {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			node := forEachNode(c, id, pre, post)
			if node != nil {
				found = true
				return node
			}
		}
	} else {
		return n
	}

	if post != nil {
		found = post(n, id)
	}

	if found {
		return n
	} else {
		return nil
	}
}

func findElementByID(n *html.Node, id *string) bool {
	if n.Type != html.ElementNode {
		return false
	}
	for _, a := range n.Attr {
		if a.Key == "id" || a.Key == "ID" {
			return a.Val == *id
		}
	}
	return false
}
