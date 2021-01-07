package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt" //go get golang.org/x/crypto/bcrypt
)

func main() {
	pwd := "password123"
	fmt.Println("Original :", pwd)
	a, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Encrypted:", string(a))
	err = bcrypt.CompareHashAndPassword(a, []byte(pwd))
	fmt.Println("Correct:", err)
	err = bcrypt.CompareHashAndPassword(a, []byte(pwd+"13212"))
	fmt.Println("Wrong  :", err)
}
