package main

// Exercis e 1.8: Modif y fetch to add the prefix http:// to each argument URL if it is missing.
// You might want to use strings.HasPrefix.

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

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
