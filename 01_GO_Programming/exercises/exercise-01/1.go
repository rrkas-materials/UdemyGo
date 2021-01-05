package main

import (
	"fmt"
)

/*
	Question:
	1	Using short declaration operator, ASSIGN these values to VARIABLES with the
		IDENTIFIERS "x", "y" and "z"
		a	42
		b	"James Bond"
		c	true
	2	Now print the values stored in those variables using
		a	single print statement
		b	multiple print statements
*/

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
