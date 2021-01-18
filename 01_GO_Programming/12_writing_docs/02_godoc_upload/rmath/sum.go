// Package rmath provides basic maths operations as an example for uploading in godoc.org
package rmath

import "fmt"

// SumAll adds unlimited numbers
func SumAll(v ...int) {
	sum := 0
	for _, val := range v {
		sum += val
	}
	fmt.Println(sum)
}
