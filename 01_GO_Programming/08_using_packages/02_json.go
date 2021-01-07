package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname    string  //field name in json is same as field name in struct
	Lname    string  `json:"lname"`         //<filed name> <type> <field name to appear in json, called tag>, changing field name in json
	Age      int     `json:"age,omitempty"` //omitempty : omits the field if its empty
	Weight   float64 `json:",omitempty"`    //keep field name same but omit if value is empty
	Password string  `json:"-"`             //ignored while processing
	Random   string  `json:"-,"`            //field name appears as "-"
}

//same as person, with some changes
type person2 struct {
	Fname    string
	Lname    string  `json:"lname"`
	Age      int     `json:"age,omitempty"`
	Weight   float64 `json:",omitempty"`
	Password string  `json:"-"`
	// Random   string  `json:"-,"` //value of "-" key in json will be ignored
	Rnd string //no such key in json, default value "" is assigned
}

//field names should start with uppercase to be read by external file

func main() {
	p1 := person{
		Fname:    "Rohnak",
		Lname:    "Agarwal",
		Age:      20,
		Weight:   65.5,
		Password: "123456",
		Random:   "kjhb",
	}
	p2 := person{
		Fname: "R",
		Lname: "A",
		Age:   21,
	}
	p3 := person{
		Fname: "R",
		Lname: "A",
	}

	//Marshalling
	people := []person{p1, p2, p3}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	str := string(bs)
	fmt.Printf("%v\n", str)
	fmt.Printf("%T\t%T\n", str, []byte(str))

	//Unmarshalling
	bs2 := []byte(str)
	var people2 []person2
	err = json.Unmarshal(bs2, &people2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people2)
}
