package main

import (
	"fmt"
)

func main() {
	foo1()
	foo2("Rohnak")
	a := foo3(4, 5)
	fmt.Println(a)
	s, d, p, q := foo4(10, 5)
	fmt.Println(s, d, p, q)
}

//simple function
func foo1() {
	fmt.Println("Hello world")
}

//with paramenter
func foo2(name string) {
	fmt.Println("Hello", name)
}

//return
func foo3(a, b int) int {
	return a + b
}

//muliple returns
func foo4(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}
