package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	default:
		io.WriteString(w, "current path: "+req.URL.Path+"<br><b>Unknown Route</b>")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
