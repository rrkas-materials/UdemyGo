package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s := "Hello"
	fmt.Println(s)
	fmt.Fprintln(os.Stdout, s)
	io.WriteString(os.Stdout, s+"\n")
}
