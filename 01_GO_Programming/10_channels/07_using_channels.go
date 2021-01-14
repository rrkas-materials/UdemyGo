package main

import "fmt"

func main() {

	c := make(chan int)

	//send
	go f1(c)

	//receive
	f2(c)

	fmt.Println("Exit")
}

//send
func f1(c chan<- int) {
	c <- 42
}

//receive
func f2(c <-chan int) {
	fmt.Println(<-c)
}
