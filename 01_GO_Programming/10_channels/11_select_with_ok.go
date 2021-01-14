package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c

	fmt.Println(v, ok)

	go func() {
		c <- 10
	}()

	v, ok = <-c

	fmt.Println(v, ok)

	go func() {
		close(c)
	}()

	v, ok = <-c

	fmt.Println(v, ok)
}
