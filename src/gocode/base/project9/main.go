package main

import (
	"fmt"
)

func main() {
	var c1 byte = 'a'
	var c2 byte = 'b'
	fmt.Println(c1, c2)

	fmt.Printf("c1=%c , c2=%c \n", c1, c2)
	var c3 int = '被'
	// var c3 byte = '被'
	fmt.Printf("c3=%c,c3=%d", c3, c3)
}
