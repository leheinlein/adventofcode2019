package main

import (
	"fmt"
	"strconv"
)

func main() {
	low := 387638
	high := 919123
	count := 0
	for i := low; i < high; i++ {
		str := strconv.Itoa(i)
		doubleRune := false
		increasing := true
		prevRune := '0' - 1
		for _, r := range str {
			if r == prevRune {
				doubleRune = true
			}
			if r < prevRune {
				increasing = false
				break
			}
			prevRune = r
		}
		if doubleRune && increasing {
			count++
		}
	}
	fmt.Printf("Result: %d", count)
}
