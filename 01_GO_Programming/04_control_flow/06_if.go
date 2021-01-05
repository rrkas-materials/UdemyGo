package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("true")
	}
	if false {
		fmt.Println("false")
	}
	if !true {
		fmt.Println("!true")
	}
	if !false {
		fmt.Println("!false")
	}
	if 4 > 5 {
		fmt.Println("4>5")
	}
	if 4 < 5 {
		fmt.Println("4<5")
	}
	if 2 == 2 {
		fmt.Println("2==2")
	}
	if 2 != 2 {
		fmt.Println("2!=2")
	}
	if !(2 != 2) {
		fmt.Println("!(2!=2)")
	}
	if x := 42; x == 42 {
		fmt.Println("x := 42; x == 42")
	}
	fmt.Println("Program ended")
}
