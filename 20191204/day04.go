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
		exactlyDoubleRune := false
		prevMatches := 0
		increasing := true
		prevRune := '0' - 1
		for _, r := range str {
			if r == prevRune {
				prevMatches++
			} else if r < prevRune {
				increasing = false
				break
			} else if prevMatches == 1 {
				exactlyDoubleRune = true
				prevMatches = 0
			} else {
				prevMatches = 0
			}
			prevRune = r
		}
		if prevMatches == 1 {
			exactlyDoubleRune = true
		}
		if exactlyDoubleRune && increasing {
			count++
		}
	}
	fmt.Printf("Result: %d", count)
}
