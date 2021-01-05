package main

import (
	"fmt"
)

/*
	Question:
	Using code from previous exercise,
	1	At package level scope, assign the following values to the 3 variables:
		a	for x assign 42
		b	for y assign "James Bond"
		c	for z assign true
	2	In func main:
		a	Use Sprintf to print all values to single string. assign the
			returned value to type string using short desciption operator to a
			variable with the identifier "s"
*/

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
