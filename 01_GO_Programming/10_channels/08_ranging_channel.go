package main

import "fmt"

func main() {

	c := make(chan int)

	//send
	go f(c)

	//receive
	for y := range c {
		fmt.Println(y)
	}

	fmt.Println("Exit")
}

//send
func f(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}
