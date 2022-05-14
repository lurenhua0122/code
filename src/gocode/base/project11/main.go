package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Tobe   bool       = false
	Maxint uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// var i, j int = 1, 2
	// k := 3
	// c, python, java := true, false, "none"
	fmt.Printf("Type: %T value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T value: %v\n", Maxint, Maxint)
	fmt.Printf("Type: %T value: %v\n", z, z)
}
