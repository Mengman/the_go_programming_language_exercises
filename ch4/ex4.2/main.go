package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	shaName := flag.String("sha", "SHA256", "SHA method name")
	str := flag.String("val", "", "the string to SHA")
	flag.Parse()
	switch *shaName {
	case "SHA384":
		fmt.Println(fmt.Sprintf("%X", sha512.Sum384([]byte(*str))))
	case "SHA512":
		fmt.Println(fmt.Sprintf("%X", sha512.Sum512([]byte(*str))))
	default:
		fmt.Println(fmt.Sprintf("%X", sha256.Sum256([]byte(*str))))
	}
}
