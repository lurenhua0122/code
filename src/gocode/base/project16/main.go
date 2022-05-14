package main

import (
	"fmt"
)

const Pi = 3.1415

func main() {
	const world = "世界"
	fmt.Println("hello", world)
	fmt.Println("happy", Pi, "day")
	const Truth = true
	fmt.Println("go rules?", Truth)
}
