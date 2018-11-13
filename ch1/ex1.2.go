package main
// Exercise 1.2: Modify the echo program to print the index and value of each of 
// its arguments, one per line. 
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(strconv.Itoa(i) + " " + arg)
	}
}
