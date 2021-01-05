package main

import (
	"fmt"
)

var a int

type rohnak int

var b rohnak

func main() {
	a = 42
	b = 50
	fmt.Println(a, b)
	fmt.Printf("%T\t%T\n", a, b)
	//a = b //error:  cannot use b (type rohnak) as type int in assignment
}
