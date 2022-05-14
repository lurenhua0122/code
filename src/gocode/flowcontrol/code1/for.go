package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	isum := 1
	for isum < 10 {
		isum += isum
	}
	fmt.Println(isum)
}
