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

	immMode = 1
)

func main() {
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
		doInput(nums, 1)
		nextIndex = index + 2
	case output:
		fmt.Println(nums[nums[index+1]])
		nextIndex = index + 2
	case end:
		nextIndex = 0
	default:
		nextIndex = -1
	}
	return
}

func doInput(nums []int, input int) {
	// "After providing 1 to the only input instruction" -> input instruction is always the first, at least for now
	nums[nums[1]] = input
}

func binaryOp(nums []int, index int, fn func(int, int) int) {
	posMode := getModes(nums[index], 3)
	var i1, i2 int
	if posMode[0] == immMode {
		i1 = nums[index+1]
	} else {
		i1 = nums[nums[index+1]]
	}
	if posMode[1] == immMode {
		i2 = nums[index+2]
	} else {
		i2 = nums[nums[index+2]]
	}
	nums[nums[index+3]] = fn(i1, i2)
}

func doMult(nums []int, index int) {
	binaryOp(nums, index, func(i1 int, i2 int) int { return i1 * i2 })
}

func doAdd(nums []int, index int) {
	binaryOp(nums, index, func(i1 int, i2 int) int { return i1 + i2 })
}

func runProgram(nums []int) error {
	var idx int
	for true {
		newIdx := processCommand(nums, idx)
		if newIdx == -1 {
			return fmt.Errorf("problem with input at %d", idx)
		}
		if newIdx == 0 {
			return nil
		} else {
			idx = newIdx
		}
	}
	return errors.New("why am I here")
}

func getModes(instruction int, numModes int) []int {
	modes := make([]int, numModes)
	modesStr := strconv.Itoa(instruction)
	for i := 0; i < numModes; i++ {
		pos := len(modesStr) - i - 3
		if pos < 0 {
			modes[i] = 0
		} else {
			m := modesStr[pos]
			if m == '1' {
				modes[i] = 1
			} else if m == '0' {
				modes[i] = 0
			} else {
				panic(fmt.Sprintf("Invalid mode: instruction %d, numMdes %d, mode %d", instruction, numModes, i))
			}
		}
	}
	return modes
}
