package main

import (
	"fmt"
)

type person struct {
	fname, lname string
	age          int
}

type student struct {
	person
	roll int
}

//method of person
func (p person) speak() {
	fmt.Println("my age:", p.age)
}

//method of student
func (s student) details() {
	fmt.Println("my name:", s.fname, s.lname)
	fmt.Println("my age:", s.age)
	fmt.Println("my roll:", s.roll)
}

//NOTE: a method can have only 1 reciever

func main() {
	s1 := student{
		person: person{
			fname: "Rohnak",
			lname: "Agarwal",
			age:   20,
		},
		roll: 25,
	}
	s1.speak()
	s1.details()
}

func foo() {
	fmt.Println("foo1")
}
