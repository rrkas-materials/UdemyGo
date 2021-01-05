package main

import (
	"fmt"
)

var a = 20 //no error, even no error if a is not used, global
var c int  //declaration only, default = 0

/*
	declare + assign = initialization

	default values:
	int			0
	boolean		false
	float		0.0
	string		""
	pointer		nil
	interface	nil
	function	nil
	slice		nil
	channel		nil
	map			nil
*/

//b:= 30	//syntax error: non-declaration statement outside function body

func main() {
	x := 42 //declare & assign (short decoration operator), local
	fmt.Println(x)
	var y = 43 //local
	fmt.Println(y)
	fmt.Println(a)
	foo()
	fmt.Println(c)
	var d int //local declaration, default value = 0
	fmt.Println(d)
}

func foo() {
	fmt.Println("again", a)
}
