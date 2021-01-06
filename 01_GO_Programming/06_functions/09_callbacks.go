package main

import (
	"fmt"
)

func main() {
	i := []int{2, 3, 4, 5}
	fmt.Println(i)
	fmt.Println("s  =", sum(i...))
	fmt.Println("es =", evensum(sum, i...))
}

func sum(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}

//f is a function with signature "(x ...int) int"
//v is list of all values from which only even are to be summed
func evensum(f func(x ...int) int, v ...int) int {
	var y []int
	for _, val := range v {
		if val%2 == 0 {
			y = append(y, val)
		}
	}
	//return sum of even numbers
	return sum(y...)
}
