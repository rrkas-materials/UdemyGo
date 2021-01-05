package main

import (
	"fmt"
)

func main() {

	if x := 40; x == 42 {
		fmt.Println("42")
	} else if y := x - 1; y == 41 {
		fmt.Println("41")
	} else if x == 40 {
		fmt.Println("40")
	} else {
		fmt.Println("else")
	}
	fmt.Println("Program ended")
}
