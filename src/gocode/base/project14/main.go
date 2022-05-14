package main

import (
	"fmt"
	"unsafe"
)

func main() {
	v := 43
	fmt.Printf("v is of type %T\n", v)

	var b = false
	fmt.Println("b=", b)
	fmt.Println("b sizeof : ", unsafe.Sizeof(b))
}
