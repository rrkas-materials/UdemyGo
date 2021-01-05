package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	i = 0
	for {
		if i > 2 {
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("Loop ended")
}
