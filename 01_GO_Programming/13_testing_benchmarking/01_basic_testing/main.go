package main

import "fmt"

func main() {
	fmt.Println(mySum(2, 3))
}

func mySum(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

/*
	command to run main.go:
	$ go run main.go
	command to run main_test.go:
	$ go test

*/
