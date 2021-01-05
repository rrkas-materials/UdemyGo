package main

import (
	"fmt"
)

/*
	Question:
	1	Create your own type. Have the underlying type be an int.
	2	create a variable of your new type with the identifier "x" using "var"
		keyword
	3	in func main:
		a	print value of variable "x"
		b	print type of variable "x"
		c	assign 42 to variable "x" using "=" operator
		d	print value of variable "x"
*/

type rohnak int

func main() {
	var x rohnak
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
