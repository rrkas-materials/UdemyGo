package main

import (
	"fmt"
)

func main() {
	x := []int{4, 6, 10, 42}
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Println(x[2]) //accessing elements by index
	fmt.Println("length = ", len(x))
	fmt.Println("capacity = ", cap(x))

	//accessing elements using loop
	for i, v := range x {
		fmt.Println(i, v)
	}

	//accessing elements using loop (alternate)
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	//slicing a slice (like python)
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[:2]) //index 2 is not included
	fmt.Println(x[1:3])

	//appending a slice
	x = append(x, 50, 100, 1024)
	fmt.Println(x)
	y := []int{1, 2, 0}
	x = append(x, y...)
	fmt.Println(x)

	//deleting from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	//making slice
	x = make([]int, 5, 10)
	fmt.Println(x)
	fmt.Println("length = ", len(x))
	fmt.Println("capacity = ", cap(x))
	x = append(x, 50, 100, 1024)
	fmt.Println(x)
	fmt.Println("length = ", len(x))
	fmt.Println("capacity = ", cap(x))
	x = append(x, 50, 100, 1024)
	fmt.Println(x)
	fmt.Println("length = ", len(x))
	fmt.Println("capacity = ", cap(x))
	//NOTE: capacity doubles when length exceeds the capacity
	// x[9] = 10 //runtime error: index out of range [9] with length 8

	//multi-dimensional
	a := []string{"James", "Bond", "Chocolate", "martini"}
	b := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	c := [][]string{a, b}
	fmt.Println(c)
	fmt.Println(c[0])
	fmt.Println(c[0][0])
	fmt.Println(c[1][1])
}
