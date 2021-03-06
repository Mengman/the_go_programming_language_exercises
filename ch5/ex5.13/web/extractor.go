package web

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

func Extract(url string, savePage bool, savePath string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	body := newRewindReader(resp.Body)
	defer body.Close()

	if strings.HasPrefix(url, "http://") {
		url = url[7:]
	} else if strings.HasPrefix(url, "https://") {
		url = url[8:]
	}
	domains := strings.Split(url, "/")
	firstDomain := domains[0]
	fileName := strings.Join(domains, "_") + ".html"
	if savePage {
		out, err := os.Create(filepath.Join(savePath, fileName))
		if err != nil {
			return nil, err
		}
		defer out.Close()
		_, err = io.Copy(out, body)
	}
	body.rewind()
	if err != nil {
		fmt.Printf(err.Error())
	}

	doc, err := html.Parse(body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if strings.Contains(link.String(), firstDomain) {
					links = append(links, link.String())
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
