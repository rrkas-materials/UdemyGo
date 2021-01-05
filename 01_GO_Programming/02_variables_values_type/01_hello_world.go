package main

import (
	"fmt"
)

//entry point of every Go program
func main() {
	fmt.Println("Hello World")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("End of program")
}

func foo() {
	fmt.Println("Foo")
}
