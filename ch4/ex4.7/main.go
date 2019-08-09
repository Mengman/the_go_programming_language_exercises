package main

import (
	"fmt"
	"unicode/utf8"
)

func rev(bytes []byte) {
	size := len(bytes)
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[size-1-i] = bytes[size-1-i], bytes[i]
	}
}

func revUTF8(bytes []byte) []byte {
	for i := 0; i < len(bytes); {
		_, size := utf8.DecodeRune(bytes[i:])
		rev(bytes[i : i+size])
		i += size
	}
	rev(bytes)
	return bytes
}

func main() {
	s := []byte("我 会写 Go 语言")
	revUTF8([]byte(s))
	fmt.Printf("%v\n", string(s))
}
