package main

import (
	"fmt"
)

func main() {
	var name string
	var age float32
	var ispass bool
	var sex byte

	fmt.Println("please input name age ispass sex")

	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Scanln(&ispass)
	fmt.Scanln(&sex)

	fmt.Printf("%s , %f , %t , %v", name, age, ispass, sex)
}
