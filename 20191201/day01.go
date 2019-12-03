package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	testMap1 := map[float64]int{
		12.0:     2,
		14.0:     2,
		1969.0:   654,
		100756.0: 33583,
	}
	testMap2 := map[float64]int{
		14.0:     2,
		1969.0:   966,
		100756.0: 50346,
	}
	for k, v := range testMap1 {
		if result := calcFuelSimple(k); result != v {
			fmt.Printf("For %f expected %d but got %d", k, v, result)
		}
	}

	for k, v := range testMap2 {
		if result := calcFuelComplex(k); result != v {
			fmt.Printf("For %f expected %d but got %d", k, v, result)
		}
	}

	file, err := os.Open("/Users/eheinlein/go/src/adventOfCode2019/20191201/day1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		num, err := strconv.ParseFloat(txt, 64)
		if err != nil {
			panic(err)
		}
		total += calcFuelComplex(num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ANSWER: %d\n", total)

}

func calcFuelSimple(num float64) int {
	f := num / 3
	i := math.Floor(f)
	return int(i) - 2
}

func calcFuelComplex(num float64) int {
	var fuelNeeded int
	var done bool
	latestFuel := num
	for !done {
		f := calcFuelSimple(latestFuel)
		if f >= 0 {
			fuelNeeded += f
			latestFuel = float64(f)
		} else {
			done = true
		}
	}
	return fuelNeeded

}
