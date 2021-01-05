package main

import (
	"fmt"
)

/*
	Question:
	1	Using the following operators, write expressions and assign their values to variables:
		==
		<=
		>=
		!=
		<
		>
	2	Now print each of the variables.
*/

func main() {
	a := 10 == 10
	b := 10 <= 10
	c := 10 >= 10
	d := 10 != 10
	e := 10 < 10
	f := 10 > 10
	fmt.Println(a, b, c, d, e, f)
}
