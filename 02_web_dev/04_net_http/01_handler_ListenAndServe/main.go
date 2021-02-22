package main

import (
	"fmt"
	"net/http"
)

type rohnak int

func (r rohnak) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("----------------------------------")
	fmt.Println(w)
	fmt.Println("----------------------------------")
	fmt.Println(req)
	val := fmt.Sprint("Value of r: ", r)
	fmt.Fprintln(w, val)
}

func main() {
	var r rohnak = 20
	http.ListenAndServe(":8080", r)
}
