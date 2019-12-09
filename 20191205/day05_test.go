package main

import (
	"testing"
)

func Test_getModes(t *testing.T) {
	cases := map[int][]int{
		1002:  {0, 1, 0},
		11101: {1, 1, 1},
		1:     {0, 0, 0},
	}
	for k, v := range cases {
		if !Equal(getModes(k, 3), v) {
			t.Errorf("Invalid mode: %d, %v", k, getModes(k, 3))
		}
	}
}

// From https://yourbasic.org/golang/compare-slices/
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
