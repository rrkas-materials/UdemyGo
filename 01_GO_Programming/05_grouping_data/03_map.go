package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"James":  32,
		"Rohnak": 20,
	}
	fmt.Println(m)
	fmt.Println(m["James"]) //key present in map
	fmt.Println(m["james"]) //no such key in map
	v, ok := m["james"]
	fmt.Println(v, ok)

	if v, ok := m["james"]; ok {
		fmt.Println("present james", v)
	}
	if v, ok := m["James"]; ok {
		fmt.Println("present James", v)
	}

	//adding elements
	m["james"] = 55
	if v, ok := m["james"]; ok {
		fmt.Println("present james", v)
	}

	//looping in map
	for k, v := range m {
		fmt.Println(k, v)
	}

	//deleting element
	delete(m, "James")
	fmt.Println(m)
	delete(m, "abc")
	fmt.Println(m)
}
