package main

import "fmt"

func main() {
	a := 20
	fmt.Println("init:", a)
	foo1(a)
	fmt.Println("foo1:", a)
	foo2(&a)
	fmt.Println("foo2:", a)
}

func foo1(a int) {
	a++
}

func foo2(a *int) {
	(*a)++
}
