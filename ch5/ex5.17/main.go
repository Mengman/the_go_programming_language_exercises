package main

// Exercise 5.17: Write a variadic function ElementsByTagName that, given an HTML node tree
// and zero or more names, returns all the elements that match one of those names. Here are two
// example calls:
// 					func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
// 					images := ElementsByTagName(doc, "img")
// 					headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RewindReader struct {
	io.ReadCloser
	content []byte
}

func newRewindReader(reader io.ReadCloser) (r RewindReader) {
	r.content, _ = ioutil.ReadAll(reader)
	r.rewind()
	return
}

func (r *RewindReader) rewind() {
	r.ReadCloser = ioutil.NopCloser(bytes.NewBuffer(r.content))
}

//ElementsByTagName implement a non-recursion version breath first search 
func ElementsByTagName(doc *html.Node, names ...string) (nodes []*html.Node) {
	visited := make(map[*html.Node]struct{})
	var queue []*html.Node

	queue = append(queue, doc)

	for len(queue) > 0 {
		node := queue[0]
		queue[0] = nil
		queue = queue[1:]

		if node.Type == html.ElementNode {
			matched := false
			for _, name := range names {
				if node.Data == name {
					matched = true
				}
			}

			if matched {
				nodes = append(nodes, node)
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if _, ok := visited[c]; !ok {
				visited[c] = struct{}{}
				queue = append(queue, c)
			}
		}
	}

	return
}

func getDoc(url string) (doc *html.Node) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("get page %s failed, status code %v\n", url, resp.StatusCode)
		return
	}

	body := newRewindReader(resp.Body)
	defer body.Close()

	doc, err = html.Parse(body)
	if err != nil {
		log.Printf("html.Parse response body failed %v", err)
		return nil
	}

	return
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run main.go https://wwww.example.com h1 h2 h3 ...")
	}

	url := os.Args[1]

	tagNames := os.Args[2:]

	doc := getDoc(url)

	if doc == nil {
		log.Println("getDoc failed")
		os.Exit(1)
	}

	nodes := ElementsByTagName(doc, tagNames...)

	fmt.Printf("found %d %v\n", len(nodes), tagNames)
}
