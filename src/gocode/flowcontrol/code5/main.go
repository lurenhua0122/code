package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("when's saturday")
	today := time.Now()

	switch {
	case today.Hour() < 12:
		fmt.Println("morning")
	case today.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("Good evening.")
	}
}
