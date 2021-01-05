package main

import (
	"fmt"
)

/*
	Question:
	1	At package level scope, using "var" keyword, create a variable with the
		identifier "y". the variable should be of the underlying type of your custom
		type "x"
	2	in func main,
		a	this should already be done:
			i	print value of variable "x"
			ii	print type of variable "x"
			iii	assign your own value to variable "x" using "=" operator
			iv	print value of variable "x"
		b	now do this:
			i	use type conversion to convert type of value stored in "x"
				to underlying type:
			ii	use assignment operator to assign that value to "y"
			iii	print value of "y"
			iv	print type of "y"
*/

type rohnak int

var y int

func main() {
	var x rohnak
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
