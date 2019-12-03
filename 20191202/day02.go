package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	add  = 1
	mult = 2
	end  = 99
)

func main() {
	input1 := "1,1,1,4,99,5,6,0,99"
	nums := parseInput(input1)
	runProgram(nums)
	fmt.Printf("%v", nums)

	file, err := os.Open("/Users/eheinlein/go/src/adventOfCode2019/20191202/day2-input.txt")
	if err != nil {
		log.Fatal(err)
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
	noun, verb := findInputs(actualInput)
	fmt.Printf("%d", 100*noun+verb)

}

func findInputs(input []int) (int, int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Printf("%d, %d\n", i, j)
			test := make([]int, len(input))
			copy(test, input)
			test[1] = i
			test[2] = j
			err := runProgram(test)
			if err != nil {
				continue
			}
			if test[0] == 19690720 {
				return i, j
			}

		}
	}
	panic("Couldn't find appropriate inputs.")
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

func processCommand(nums []int, index int) (bool, error) {
	if nums[index] == end {
		return true, nil
	}
	op1 := nums[nums[index+1]]
	op2 := nums[nums[index+2]]
	output := nums[index+3]
	switch nums[index] {
	case add:
		nums[output] = op1 + op2
		return false, nil
	case mult:
		nums[output] = op1 * op2
		return false, nil
	default:
		return false, errors.New("Invalid command: " + strconv.Itoa(nums[index]))
	}
}

func runProgram(nums []int) error {
	var done bool
	var idx int
	var err error
	for true {
		done, err = processCommand(nums, idx)
		if err != nil {
			return err
		}
		if !done {
			idx += 4
		} else {
			return nil
		}
	}
	return errors.New("why am I here")
}
