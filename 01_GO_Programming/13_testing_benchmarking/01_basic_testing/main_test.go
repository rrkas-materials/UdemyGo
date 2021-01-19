package main

import "testing"

func TestMySum(t *testing.T) {
	s := mySum(2, 3)
	if s != 5 {
		t.Error("expected 5, got", s)
	}
}

func TestSum(t *testing.T) {
	s := mySum(2, 3)
	if s != 5 {
		t.Error("expected 5, got", s)
	}
}

//Note: both functions will work same
