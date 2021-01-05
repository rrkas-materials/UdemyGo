package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e
	f
)

func main() {
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Printf("%v\t%T\n", d, d)
	fmt.Printf("%v\t%T\n", e, e)
	fmt.Printf("%v\t%T\n", f, f)
}
