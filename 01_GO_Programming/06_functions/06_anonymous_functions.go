package main

import (
	"fmt"
)

func main() {

	func(s string) {
		fmt.Println("Hello", s)
	}("abc")

	func(a ...int) {
		s := 0
		for _, v := range a {
			s += v
		}
		fmt.Println("sum:", s)
	}(1, 2, 3, 4, 5, 7, 8, 9)

	a := func(a ...int) int {
		s := 0
		for _, v := range a {
			s += v
		}
		return s
	}(1, 2, 3, 4, 5, 7, 8, 9)
	fmt.Println("sum:", a)
}
