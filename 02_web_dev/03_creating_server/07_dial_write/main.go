package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}

/*
	terminal 1:
		$ cd 02_web_dev/03_creating_server/03_read/
		$ go run main.go
	terminal 2:
		$ cd 02_web_dev/03_creating_server/07_dial_write/
		$ go run main.go
*/
