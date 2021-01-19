package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{10, 10}, 20},
		test{[]int{5, 15}, 20},
		test{[]int{19, 1}, 20},
		test{[]int{10, 10}, 20},
	}

	for _, v := range tests {
		s := mySum(v.data...) + 1
		if s != v.answer {
			t.Error("expected", v.answer, "got", s)
		}
	}
}
