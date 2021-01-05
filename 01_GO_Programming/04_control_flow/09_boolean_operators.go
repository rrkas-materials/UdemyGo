package main

import (
	"fmt"
)

/*
	&& :  	and
	|| :	or
	!  :	not
*/

func main() {
	fmt.Println("And-------------------------")
	fmt.Println("true && true \t= ", true && true)
	fmt.Println("true && false \t= ", true && false)
	fmt.Println("false && true \t= ", false && true)
	fmt.Println("false && false \t= ", false && false)
	fmt.Println("OR-------------------------")
	fmt.Println("true || true \t= ", true || true)
	fmt.Println("true || false \t= ", true || false)
	fmt.Println("false || true \t= ", false || true)
	fmt.Println("false || false \t= ", false || false)
	fmt.Println("Not-------------------------")
	fmt.Println("!true \t\t\t= ", !true)
	fmt.Println("!false \t\t\t= ", !false)
}
