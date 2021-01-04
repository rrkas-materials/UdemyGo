package main

import (
	"fmt"
)

/*
	Question:
	Write a program that prints a number in ​ decimal, binary, and hex
*/

func main() {
	a := 10
	fmt.Printf("%d\t\t%b\t\t%#x", a, a, a)
}
