package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello")
	fmt.Println(n, err)
	n0, _ := fmt.Println("123")
	fmt.Println(n0)
}
