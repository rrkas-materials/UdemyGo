package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println(x)
	}
	g(20)
}
