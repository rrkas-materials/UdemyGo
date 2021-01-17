package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		f(i)
	}
	fmt.Println("Returned normally from f.")
}

func f(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	g(i, i+3)
}

func g(i, n int) {
	if i > n {
		fmt.Println("Panicking!", i)
		panic(fmt.Sprintf("%v", i))
	}
	g(i+1, n)
}
