package main

import (
	"fmt"
)

/*
	Question:
	1	Write a program that
		a	assigns an int to a variable
		b	prints that int in decimal, binary, and hex
		c	shifts the bits ​ of that int over 1 position to the left, and assigns that to a variable
		d	prints that variable in decimal, binary, and hex
*/

func main() {
	a := 10
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}
