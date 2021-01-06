package main

import (
	"fmt"
)

func main() {
	fmt.Println(f1())
	fmt.Println(f2(1, 2)(3)) //f2(1,2) is a function, not a value. f2(1,2)(3) is a value
}

func f1() string {
	return "hello"
}

func f2(a, b int) func(int) string {
	return func(c int) string {
		return fmt.Sprint(a, b, c)
	}
}
