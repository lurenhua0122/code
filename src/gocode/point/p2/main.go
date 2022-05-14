package main

import (
	"fmt"
)

func main() {
	var ii int = 12
	fmt.Println("address:", &ii)

	var ptr *int = &ii
	fmt.Println(*ptr)
}
