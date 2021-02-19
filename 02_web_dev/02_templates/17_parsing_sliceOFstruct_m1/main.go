package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type person struct {
	Fname, Lname string
	ID           int
}

func main() {
	d1 := person{"A", "B", 30}
	d2 := person{"C", "d", 32}
	d3 := person{"E", "f", 35}
	d4 := person{"G", "h", 38}
	data := []person{d1, d2, d3, d4}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}

/*
	$ go run main.go > index.html
*/
