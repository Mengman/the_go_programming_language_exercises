package main

// Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but
// supports a command-line flag to print the SHA384 or SHA512 hash instead.

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	shaName := flag.String("sha", "SHA256", "SHA method name")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		switch *shaName {
		case "SHA384":
			fmt.Println(fmt.Sprintf("%X", sha512.Sum384([]byte(in))))
		case "SHA512":
			fmt.Println(fmt.Sprintf("%X", sha512.Sum512([]byte(in))))
		default:
			fmt.Println(fmt.Sprintf("%X", sha256.Sum256([]byte(in))))
		}
	}

}
