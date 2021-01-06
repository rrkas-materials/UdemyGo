package main

import (
	"fmt"
)

/*
	NOTE:

	func foo(b ...int, a string) //error
	func foo(a string, b ...int) //no error

	variadic parameter should be final parameter

*/

func main() {
	fmt.Println(foo(2, 3, 4, 5, 6))
	a := []int{1, 2, 3}
	fmt.Println(foo(a...))
	fmt.Println(foo())
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}
