// Package a provides ACME inc math solutions
package a

import "fmt"

// Sum adds unlimited numbers
func Sum(v ...int) {
	sum := 0
	for _, val := range v {
		sum += val
	}
	fmt.Println(sum)
}
