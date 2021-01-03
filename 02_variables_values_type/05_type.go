package main

import (
	"fmt"
)

//variable can hold value STATICLY

var y = 42
var z string = "Rohnak Agarwal"

func main() {
	fmt.Print(y, " ")
	fmt.Printf("%T\n", y)
	fmt.Print(z, " ")
	fmt.Printf("%T\n", z)
	//z = 50 //illegal, cannot use 50 (type untyped int) as type string in assignment
	fmt.Print(z, " ")
	fmt.Printf("%T\n", z)
	rawString := `I said, "Hello!"`
	interpretedString := "abc"
	fmt.Println(rawString, interpretedString)
}
