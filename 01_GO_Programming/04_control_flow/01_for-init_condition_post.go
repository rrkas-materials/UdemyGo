package main

import (
	"fmt"
)

/*
	syntax:
	for initialiation; condition; post {
		statements
	}
*/

func main() {
	fmt.Println("Loop started")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	j := 0
	for ; j < 5; j++ { //init is optional, same for condition and post
		fmt.Println(j)
	}
	fmt.Println("Loop ended")
}
