package main

import (
	"fmt"
)

var y string
var z int

func main() {

	fmt.Println("y=", y, ".")
	fmt.Printf("%T\n", y)
	y = "a"
	fmt.Println("y=", y, ".")
	fmt.Printf("%T\n", y)

	fmt.Println("z=", z, ".")
	fmt.Printf("%T\n", z)
}
