package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type rohnak int

func (r rohnak) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(req.Form)
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var r rohnak
	http.ListenAndServe(":8080", r)
}
