package main

// Exercise 5.19: Use panic and recover to write a function that contains no return statement
// yet returns a non-zero value.

import(
	"fmt"
)

func foo() (s string) {
	defer func(){
		recover()
		s = "defer"
	}()

	panic("foo")
}

func main() {
	fmt.Println(foo())
}