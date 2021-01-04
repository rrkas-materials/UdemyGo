package main

import (
	"fmt"
)

/*
	Question:
	1	Using ​ iota​ , create 4 constants for the next 4 years.
	2 	Print the constant values.
*/

const (
	_     = iota
	next1 = 2020 + iota
	next2 = 2020 + iota
	next3 = 2020 + iota
	next4 = 2020 + iota
)

func main() {
	fmt.Println("", next1, "\n", next2, "\n", next3, "\n", next4)
}
