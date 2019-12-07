package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	add    = 1
	mult   = 2
	input  = 3
	output = 4
	end    = 99

	posMode = 0
	immMode = 1
)

func main() {
	input1 := "1,1,1,4,99,5,6,0,99"
	nums := parseInput(input1)
	runProgram(nums)
	fmt.Printf("%v", nums)

	file, err := os.Open("/Users/eheinlein/go/src/adventOfCode2019/20191205/day5-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var actualInput []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		actualInput = parseInput(scanner.Text())
	}
	if actualInput == nil {
		panic("Nil input")
	}
	err = runProgram(actualInput)
	if err != nil {
		panic(err)
	}

}

func parseInput(s string) []int {
	splitStrings := strings.Split(s, ",")
	nums := make([]int, len(splitStrings))
	var err error
	for i, v := range splitStrings {
		nums[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	return nums
}

func processCommand(nums []int, index int) (nextIndex int) {
	switch nums[index] % 100 {
	case add:
		doAdd(nums, index)
		nextIndex = index + 4
	case mult:
		doMult(nums, index)
		nextIndex = index + 4
	case input:
		doInput(nums, index)
		nextIndex = index + 2
	case output:
		fmt.Println(nums[index+1])
		nextIndex = index + 2
	case end:
		nextIndex = 0
	default:
		nextIndex = -1
	}
	return
}

func doInput(nums []int, index int) {

}

func doMult(nums []int, index int) {

}

func doAdd(nums []int, index int) {

}

func runProgram(nums []int) error {
	var idx int
	for true {
		newIdx := processCommand(nums, idx)
		if newIdx == -1 {
			return fmt.Errorf("Problem with input at %d", idx)
		}
		if newIdx == 0 {
			return nil
		} else {
			idx = newIdx
		}
	}
	return errors.New("why am I here")
}
