package main

import (
	"errors"
	"fmt"
	"log"
)

type latlng struct {
	lat, lng float64
	err      error
}

func (l latlng) Error() string {
	return fmt.Sprintln(l.lat, l.lng, l.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Printf("%T\n", err)
		log.Println(err)
	}

	_, err = sqrt2(-10)
	if err != nil {
		log.Printf("%T\n", err)
		log.Println(err)
	}

	_, err = sqrt3(-10)
	if err != nil {
		log.Printf("%T\n", err)
		log.Println(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New(fmt.Sprint("math error: square root of negative number: ", f))
	}
	return 1, nil
}

func sqrt2(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math error: square root of negative number: %v", f)
	}
	return 1, nil
}

func sqrt3(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("math error: square root of negative number: %v", f)
		return 0, latlng{2, 3, err}
	}
	return 1, nil
}
