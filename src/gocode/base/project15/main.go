package main

import (
	"fmt"
)

func main() {
	v := 43
	i := v
	f := 3.1415
	g := 0.876 + 0.5i
	fmt.Printf("v is of type %T %T %T %T\n", v, i, f, g)
}
