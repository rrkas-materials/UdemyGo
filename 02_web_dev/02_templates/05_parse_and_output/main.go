package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(1, err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(2, err)
	}
	defer nf.Close()
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(3, err)
	}
}

/*
	$ go run main.go
*/
