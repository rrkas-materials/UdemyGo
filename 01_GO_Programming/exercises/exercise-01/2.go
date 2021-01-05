package main

import (
	"fmt"
)

/*
	Question:
	1	Use var to declare 3 variables. These variables should have package level scope.
		Do not assign values to the variables. Use the following identifiers for the
		variables and make sure the variables store values of following type:
		a	identifier "x" type int
		b	identifier "y" type string
		c	identifier "z" type bool
	2	In func main:
		a	print out values for each identifier
*/

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
