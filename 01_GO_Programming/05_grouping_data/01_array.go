package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 20 //at position 4th
	fmt.Println(x)
	fmt.Println("len = ", len(x))
}
