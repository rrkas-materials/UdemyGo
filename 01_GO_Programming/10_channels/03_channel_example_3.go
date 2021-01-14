package main

import "fmt"

func main() {
	c := make(chan int, 1) //buffer : 1
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
	//works
}
