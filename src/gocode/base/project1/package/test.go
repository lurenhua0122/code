package main

import "fmt"

func sayOK() {
	fmt.Println("fdsfdsfsfs")
}

func cc() {
	sayOK()
	sum, sub := getsum(11, 12)
	fmt.Println("sum : ", sum, "sub: ", sub)

}

func getsum(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
