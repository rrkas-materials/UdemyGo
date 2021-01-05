package main

import (
	"fmt"
)

func main() {
	switch {
	case false, true:
		fmt.Println("false")
	case (2 == 2):
		fmt.Println("2==2")
		fallthrough
	case (3 < 4):
		fmt.Println("3<4")
		fallthrough
	case (2 > 5):
		fmt.Println("2>5")
	default:
		fmt.Println("default")
	}

	fmt.Println("---------------------------")

	a := 20
	switch a {
	case 15:
		fmt.Println("15")
	case (10 * 2):
		fmt.Println("20")
	case 25:
		fmt.Println("25")
	default:
		fmt.Println("default")
	}

	fmt.Println("---------------------------")

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	fmt.Println("Program ended")
}
