package main

import (
	"fmt"
	"runtime"
)

var a int
var b float64
var c uint

func main() {
	x := 42
	y := 42.345678
	//a = 42.1	  //error: constant 42.1 truncated to integer
	b = 42.345678
	c = 42
	fmt.Printf("%v\t\t\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Printf("%v\t\t\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t\t\t%T\n", c, c)
	a = -20
	//c = -20    //error: constant -20 overflows uint
	fmt.Printf("%v\t\t\t%T\n", a, a)
	fmt.Printf("%v\t\t\t%T\n", c, c)
	fmt.Println("GOARCH", runtime.GOARCH)
	fmt.Println("GOOS", runtime.GOOS)
	fmt.Println("GOROOT", runtime.GOROOT)
}
