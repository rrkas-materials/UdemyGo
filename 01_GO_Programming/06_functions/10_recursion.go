package main

import (
	"fmt"
)

func main() {
	fmt.Println("3! = ", fact(3))
	fmt.Println("5! = ", fact(5))

	for i := 1; i < 10; i++ {
		fmt.Println(i, ":", fibbo(i))
	}
}

func fact(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func fibbo(n int) int {
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibbo(n-1) + fibbo(n-2)
}
