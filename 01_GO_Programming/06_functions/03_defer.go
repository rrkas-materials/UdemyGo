package main

import (
	"fmt"
)

func main() {
	foo1()
	foo2()
	defer foo1()
	foo2()
}

func foo1() {
	fmt.Println("foo1")
}

func foo2() {
	fmt.Println("foo2")
}
