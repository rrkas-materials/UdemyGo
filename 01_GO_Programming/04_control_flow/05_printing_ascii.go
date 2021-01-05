package main

import (
	"fmt"
)

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t\t%x\t\t%#U\n", i, i, i)
	}
	fmt.Println("Loop ended")
}
