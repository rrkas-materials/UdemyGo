package main

import "fmt"

func splitNamed(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func splitUnnamed(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}

func main() {
	fmt.Println(splitNamed(17))
	fmt.Println(splitUnnamed(17))
}
