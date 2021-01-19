package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("Exit")
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("even", v)
		case v := <-o:
			fmt.Println("odd ", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From !ok", i, ok)
				return
			}
			fmt.Println("From ok", i, ok)
		}
	}
}
