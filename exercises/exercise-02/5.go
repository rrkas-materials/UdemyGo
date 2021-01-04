package main

import (
	"fmt"
)

/*
	Question:
	1	Create a variable of type string using a​ raw string literal​.
	2	Print it.
*/

func main() {
	a :=
		`Rohnak said, 	"this
					is
					a
					raw
					string"
	`
	fmt.Print(a)
}
