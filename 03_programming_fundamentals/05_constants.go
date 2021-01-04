package main

import (
	"fmt"
)

const hello = "Hello world"             //untyped constant
const typedHello string = "Hello World" //typed constant
const (
	a        = 42
	b        = 42.3456
	c        = "James Bond"
	d string = "123"
) //multiple const declarations

func main() {
	fmt.Println(hello, typedHello, a, b, c, d)
}
