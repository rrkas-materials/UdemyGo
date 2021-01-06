package main

import "fmt"

/*
	Method Sets:
		A type may have a method set associated with it.
		The method set of an interface type is its interface.
		The method set of any other type T consists of all methods declared with receiver type T.
		The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T
		(that is, it also contains the method set of T).
		Further rules apply to structs containing embedded fields, as described in the section on struct types.
		Any other type has an empty method set.
		In a method set, each method must have a unique non-blank method name.
	https://flaviocopes.com/golang-methods-receivers/
*/

type person struct {
	fname, lname string
	age          int
}

//non pointer receiver
func (p person) talk() {
	fmt.Println("talking :", p.fname, p.lname, p.age)
}

//pointer receiver
func (p *person) talk2() {
	fmt.Println("talking2:", p.fname, p.lname, p.age)
}

type human interface {
	talk()
	talk2()
}

func main() {
	p := person{"Rohnak", "Agarwal", 20}
	fmt.Printf("%T\n", p)
	p.talk() //works on non-pointer value
	p.talk2()
	//info(p) //cannot use p (type person) as type human in argument to info: person does not implement human (talk2 method has pointer receiver)

	p2 := &person{"R", "A", 20}
	fmt.Printf("%T\n", p2)
	p2.talk() //works on pointer value
	p2.talk2()
	info(p2)
}

func info(h human) {
	h.talk()
	h.talk2()
}
