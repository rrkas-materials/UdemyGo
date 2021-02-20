package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

/*
	terminal 1:
		$ cd 02_web_dev/03_creating_server/02_write/
		$ go run main.go
	terminal 2:
		$ cd 02_web_dev/03_creating_server/06_dial_read/
		$ go run main.go
*/
