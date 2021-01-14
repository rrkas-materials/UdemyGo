package main

import "fmt"

func main() {
	fmt.Println("c------------------")

	c := make(chan int, 1)
	c <- 41
	fmt.Printf("%T\n", c)
	fmt.Println(<-c)

	fmt.Println("cs------------------")

	cs := make(chan<- int, 1)
	cs <- 41
	fmt.Printf("%T\n", cs)
	// fmt.Println(<-cs) //invalid operation: <-c (receive from send-only type chan<- int)

	fmt.Println("cr------------------")

	cr := make(<-chan int, 1)
	// cr <- 41 //invalid operation: d <- 41 (send to receive-only type <-chan int)
	fmt.Printf("%T\n", cr)
	// fmt.Println(<-cr) // error: all goroutines are asleep - deadlock!

	fmt.Println("specific to general------------------")

	// c = cr //error: cannot use cr (type <-chan int) as type chan int in assignment
	// c = cs //error: cannot use cs (type chan<- int) as type chan int in assignment

	fmt.Println("specific to specific------------------")

	// cr = cs //error: cannot use cs (type chan<- int) as type <-chan int in assignment

	fmt.Println("general to specific------------------")

	cr = c
	cs = c

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	fmt.Println("conversion------------------")

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", (chan<- int)(c))
	fmt.Printf("%T\n", (<-chan int)(c))

}
