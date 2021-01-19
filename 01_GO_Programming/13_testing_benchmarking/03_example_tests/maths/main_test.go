package maths

import (
	"fmt"
)

func ExampleSum2() {
	fmt.Println(Sum2(2, 5))
	//Output:
	//7
}

/*
	testing a package
	$ go test

	testing the package and all its subpackages
	$ go test ./...
*/
