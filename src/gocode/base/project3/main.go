package main

import (
	"fmt"
)

var b1 = 100
var b2 = 200
var b3 = "global"

var (
	b4 = 100
	b5 = "name"
)

func main() {
	// var i1 int
	// var i2 int
	// var i3 int
	var n1, n2, n3 int
	fmt.Println("n1 n2 n3 :", n1, n2, n3)

	var s1, s2, s3 = 100, "tom", 0.3
	fmt.Println("s:", s1, s2, s3)

	c1, c2, c3 := 111, "jery", 0.3333
	fmt.Println("c:", c1, c2, c3)

	fmt.Println("b:", b1, b2, b3)
	fmt.Println("b:", b4, b5)
}
