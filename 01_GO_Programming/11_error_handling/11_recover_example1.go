package main

import (
	"log"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered:", r)
		}
	}()

	for i := 0; i < 3; i++ {
		_, err := os.Open("no-file.txt")
		if err != nil {
			log.Panicln(i, "error:", err)
		}
	}

}
