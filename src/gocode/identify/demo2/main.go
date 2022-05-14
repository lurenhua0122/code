package main

import (
	"fmt"
)

func main() {
	fmt.Println(10 / 4)
	var n1 float32 = 10 / 4
	fmt.Println(n1)

	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	fmt.Println("10%3=", 10%3)
	fmt.Println("-10%3=", -10%3)
	fmt.Println("10%-3=", 10%-3)
	fmt.Println("-10%-3=", -10%-3)

	var i int = 10
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}
