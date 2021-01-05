package main

import (
	"fmt"
)

var a int

type rohnak int

var b rohnak

//In Go, there is NO concept of type casting

func main() {
	a = 42
	b = 50
	fmt.Println(a, b)
	fmt.Printf("%T\t%T\n", a, b)
	a = int(b) //type conversion, NOT casting
	fmt.Println(a, b)
	fmt.Printf("%T\t%T\n", a, b)
}
