package main

// Exercise 7.4: The strings.NewReader function returns a value that satisfies the io.Reader
// interface (and others) by reading from its argument, a string. Implement a simple version of
// NewReader yourself, and use it to make the HTML parser (§5.2) take input fro m a string.

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

type Reader struct {
	s string
	i int64
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}

	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func NewReader(s string) *Reader {
	return &Reader{s, 0}
}

const (
	htmlPage = `<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<h1>hello world</h1>
</body>
</html>
`
)

func main() {
	doc, err := html.Parse(NewReader(htmlPage))

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

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(counter, c)
	}
}
