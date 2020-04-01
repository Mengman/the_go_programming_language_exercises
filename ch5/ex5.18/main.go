package main

// Exercise 5.18: Without changing its behavior, rewrite the fetch function to use defer to close
// the writable file.


import (
	"log"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "./index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		err = f.Close()
		log.Printf("close file in defer with error code %v\n", err)
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

func main() {
	url := os.Args[1]

	_, _, err := fetch(url)
	
	if err != nil {
		log.Printf("fetch(%s) return err %v\n", url, err)
	}
}
