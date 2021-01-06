package main

import (
	"fmt"
)

// any struct containing speak is of type human
type human interface {
	speak()
}

type person struct {
	fname, lname string
	age          int
}

type person2 struct {
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

func main() {
	p1 := person{
		fname: "Rohnak",
		lname: "Agarwal",
		age:   20,
	}
	p2 := person2{
		fname: "Rohnak",
		lname: "Agarwal",
		age:   20,
	}
	s1 := student{
		person: p1,
		roll:   25,
	}
	p1.speak()
	fmt.Println(p2)
	//p2.speak() //error: p2.speak undefined (type person2 has no field or method speak)
	s1.speak()
	s1.details()
	foo(p1) //as person contains speak(), it is of type human, hance, can be passed
	foo(s1) //as student contains speak(), it is of type human, hance, can be passed
	//foo(p2) //error: cannot use p2 (type person2) as type human in argument to foo: person2 does not implement human (missing speak method)
}

//polymorphism: same functions behaves differently when person and student are passed
func foo(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Person:", h.(person).fname) //asserting h as person
	case student:
		fmt.Println("Student:", h.(student).roll)
	}
}
