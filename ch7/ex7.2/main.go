package main

import (
	"fmt"
	"io"
	"os"
)

type CWriter struct {
	writer io.Writer
	count  int64
}

func (c *CWriter) Write(p []byte) (n int, err error) {
	count, err := c.writer.Write(p)
	c.count += int64(count)
	return count, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CWriter{
		writer: w,
	}

	return &cw, &cw.count
}

const s1 = `Do not go gentle into that good night,
Old age should burn and rave at close of day;
Rage, rage against the dying of the light.`

const s2 = `Though wise men at their end know dark is right,
Because their words had forked no lightning they
Do not go gentle into that good night.`

func main() {
	cw, num := CountingWriter(os.Stdout)

	n, err := cw.Write([]byte(s1))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nwrite %d bytes, total %d bytes\n", n, *num)

	n, err = cw.Write([]byte(s2))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nwrite %d bytes, total %d bytes\n", n, *num)
}
