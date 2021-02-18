package main

import (
	"fmt"
)

// var y int //global scope, default value = 0

type person struct {
	fname, Lname string //lower case is not visible outside package,uppercase is visible outside package
	rollno       int
}

type person2 struct {
	fname, Lname string
}

//function syntax: func (receiver) identifier(parameters) (returns) {code}
func (p person) speak() {
	fmt.Println(p.fname, `says, "Hello!"`)
}

func (p person2) speak() {
	fmt.Println(p.fname, `says, "Hello!"`)
}

type student struct {
	person     //same as inherits person
	rollno int //different than person.rollno
	redg   string
}

func (s student) studentData() {
	fmt.Println(s.redg)
}

//any struct having speak() as its member is a type of human interface
type human interface {
	speak()
}

func main() {
	// x := 7 //short variable declaration operator
	// fmt.Println(x)
	// fmt.Printf("%T\n", x)
	// fmt.Println(y)
	// y = 10
	// fmt.Println(y)

	//Composite literal syntax : datatype{data}

	// xi := []int{2, 4, 6} //slice of int
	// fmt.Println(xi)

	// m := map[string]int{
	// 	"Rohnak": 45,
	// 	"Abc":    60,
	// }
	// fmt.Println(m)

	// type hotdog int
	// var t hotdog = 7
	// fmt.Println(t)
	// fmt.Printf("%T\n", t)

	// p1 := person{"Rohnak", "Agarwal", 20}
	// p2 := person{
	// 	fname:  "Rohnak2",
	// 	rollno: 30,
	// }
	// fmt.Println(p1, p2)

	// p1.speak()
	// p2.speak()
	// s := student{p1, 30, "50"}
	// fmt.Println(s.rollno, s.person.rollno) //if rollno would have absent in student, rollno in person will level up and s.rollno = s.person.rollno
	// s.person.speak()                        //inherited members can be accessed through base class name
	// s.speak()                               //inherited members are levelled up
	// s.studentData()

	// s2 := student{
	// 	person{
	// 		"a", "b", 10,
	// 	},
	// 	20,
	// 	"30",
	// }
	// fmt.Println(s2)

	// p1 := person{"a", "b", 20}
	// p2 := person2{"c", "d"}
	// humanSpeak(p1) //polymorphism
	// humanSpeak(p2) //polymorphism
}

func humanSpeak(h human) {
	h.speak()
}
