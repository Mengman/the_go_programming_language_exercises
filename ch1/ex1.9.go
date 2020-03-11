package main

// Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if url[:7] != prefix {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s status code %s\n", url, resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
