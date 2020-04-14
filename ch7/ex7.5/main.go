package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Exercise 7.5: The LimitReader function in the io package accepts an io.Reader r and a
// number of bytes n, and returns another Reader that reads from r but reports an end-of-file
// condition after n bytes. Implement it.
// func LimitReader(r io.Reader, n int64) io.Reader

type LReader struct {
	reader io.Reader
	read   int64
	limit  int64
}

func (r *LReader) Read(p []byte) (n int, err error) {
	if r.read >= r.limit {
		return 0, io.EOF
	}

	i := r.limit - r.read

	if i < int64(len(p)) {
		n, err = r.reader.Read(p[:i])
	} else {
		n, err = r.reader.Read(p)
	}

	r.read += i

	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LReader{
		reader: r,
		limit:  n,
	}
}

func main() {
	s := "123456789"

	r := strings.NewReader(s)

	p := make([]byte, 9)

	n, err := r.Read(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("StringReader read %d '%s'\n", n, string(p))

	lr := LimitReader(strings.NewReader(s), 4)

	p = make([]byte, 9)

	n, err = lr.Read(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("LimitReader read %d '%s'\n", n, string(p))
}
