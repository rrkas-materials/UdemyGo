package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

//can be used by sort.Strings
func (p Person) String() string {
	return fmt.Sprintf("%v: %v", p.Name, p.Age)
}

//age wise sorting
type PersonsByAge []Person

//sort.Interface implementations
func (p PersonsByAge) Len() int           { return len(p) }
func (p PersonsByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PersonsByAge) Less(i, j int) bool { return p[i].Age < p[j].Age } // the deciding condition

//---------------------------------------------------------------------------------------------------------------
//name wise sorting
type PersonsByName []Person

//sort.Interface implementations
func (p PersonsByName) Len() int           { return len(p) }
func (p PersonsByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PersonsByName) Less(i, j int) bool { return p[i].Name < p[j].Name } // the deciding condition

func main() {
	a := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	b := []string{"James", "Quality", "Madras", "Moneypenny", "Drawer"}

	fmt.Println("ints-----------------")
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)

	fmt.Println("strings--------------")
	fmt.Println(b)
	sort.Strings(b)
	fmt.Println(b)

	fmt.Println("custom---------------")
	p1 := Person{"Jame", 38}
	p2 := Person{"Bran", 22}
	p3 := Person{"Kate", 31}
	p4 := Person{"Alee", 42}
	people := []Person{p1, p2, p3, p4}
	fmt.Println("Original :\t", people)
	sort.Sort(PersonsByAge(people))
	fmt.Println("Age  wise:\t", people)
	sort.Sort(PersonsByName(people))
	fmt.Println("Name wise:\t", people)
}
