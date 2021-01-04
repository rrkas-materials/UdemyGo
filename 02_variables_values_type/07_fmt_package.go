package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Print(y, "\n")    //only prints, no new line
	fmt.Println(y)        //prints and adds a new line
	fmt.Printf("%T\n", y) //prints in a formatted way
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%v\n", y)

	fmt.Println("-----------------")

	//NOTE: the following statements return string, not print in console
	var s string
	s = fmt.Sprint(y, "\n") //only prints to a string, no new line
	fmt.Print(s)
	s = fmt.Sprintln(y) //prints to string and adds a new line
	fmt.Print(s)
	s = fmt.Sprintf("%T\n", y) //prints to string in a formatted way
	fmt.Print(s)
	s = fmt.Sprintf("%b\n", y)
	fmt.Print(s)
	s = fmt.Sprintf("%x\n", y)
	fmt.Print(s)
	s = fmt.Sprintf("%#x\n", y)
	fmt.Print(s)
}
