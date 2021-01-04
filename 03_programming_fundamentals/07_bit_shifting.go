package main

import (
	"fmt"
)

const (
	_   = iota
	kb2 = 1 << (10 * iota)
	mb2 = 1 << (10 * iota)
	gb2 = 1 << (10 * iota)
)

func main() {

	printb(1)
	printb(1 * 2)
	x := 4
	printb(x)
	y := x << 1
	printb(y)

	kb := 1024
	mb := kb * 1024
	gb := mb * 1024
	printb(kb)
	printb(mb)
	printb(gb)
	printb(kb2)
	printb(mb2)
	printb(gb2)
}

func printb(x int) {
	fmt.Printf("%10d\t\t\t%35b\n", x, x)
}
