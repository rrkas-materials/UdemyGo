package main

import (
	"fmt"
)

func main() {

	a := 42
	d := 42
	b := &a //b stores address of a
	c := *b //dereferncig b, simply put, getting "value at address stored in b"
	printdetails("a", a)
	printdetails("*&a", *&a)
	printdetails("b", b)
	printdetails("&b", &b)
	printdetails("d", d)
	printdetails("&d", &d)
	printdetails("c", c)
	*b = 50 //"value at address stored in b" is changed to 50
	printdetails("*b = 50; a", a)
	printdetails("a", a)
	printdetails("*&a", *&a)
	printdetails("b", b)
	printdetails("&b", &b)
	printdetails("c", c)
}

func printdetails(name string, a interface{}) {
	fmt.Printf("%10v\t\tValue: %20v\t\tType: %15T\n", name, a, a)
}
