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

func main() {
	alpha := []string{"A", "B", "c", "d"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", alpha)
	if err != nil {
		log.Fatal(err)
	}
}

/*
	$ go run main.go > index.html
*/
