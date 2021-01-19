package rrka

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Rohnak")
	if s != "Welcome Rohnak" {
		t.Error("got", s, "expected", "Welcome Rohnak")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Rohnak"))
	// Output:
	// Welcome Rohnak
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Rohnak")
	}
}
