package main

import (
	"fmt"
)

func main() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("%d\t%v\t%#x\t%q\n", i, v, v, v)
	}
}
