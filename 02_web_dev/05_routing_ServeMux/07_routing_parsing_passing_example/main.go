package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/dog/", http.HandlerFunc(bar)) // same as http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", rohnak)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func rohnak(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("rohnak.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "rohnak.gohtml", "Rohnak Agarwal")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
