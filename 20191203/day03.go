package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

const (
	UP    = "U"
	DOWN  = "D"
	RIGHT = "R"
	LEFT  = "L"
)

var matcher = regexp.MustCompile(`([RUDL])(\d+)`)

func main() {
	testInput := "R8,U5,L5,D3\nU7,R6,D4,L4"
	fmt.Printf("Distance: %d\n", processInput(testInput))
	testInput1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"
	fmt.Printf("Distance1: %d\n", processInput(testInput1))
	testInput2 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	fmt.Printf("Distance2: %d\n", processInput(testInput2))

	fileText, err := ioutil.ReadFile("/Users/eheinlein/go/src/adventOfCode2019/20191203/day3-input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("REAL DISTANCE: %d\n", processInput(string(fileText)))


}

func processInput(str string) (closest int) {
	inputs := strings.Split(str, "\n")
	points := processFirstWire(inputs[0])
	crossPoints := processSecondWire(inputs[1], points)
	return calClosest(crossPoints)
}

func processFirstWire(s string) map[point]bool {
	points := make(map[point]bool)
	matches := matcher.FindAllStringSubmatch(s, -1)
	currentPoint := point{}
	for _, instruction := range matches {
		moveNum, _ := strconv.Atoi(instruction[2])
		for i := 0; i < moveNum; i++ {
			newPoint := point{
				x: currentPoint.x,
				y: currentPoint.y,
			}
			switch instruction[1] {
			case UP:
				newPoint.y += 1
			case DOWN:
				newPoint.y -= 1
			case RIGHT:
				newPoint.x += 1
			case LEFT:
				newPoint.x -= 1
			}
			points[newPoint] = true
			currentPoint = newPoint
		}

	}
	return points
}

func processSecondWire(s string, wire1Points map[point]bool) []point {
	crossPoints := make([]point, 0)
	matches := matcher.FindAllStringSubmatch(s, -1)
	currentPoint := point{}
	for _, instruction := range matches {
		moveNum, _ := strconv.Atoi(instruction[2])
		for i := 0; i < moveNum; i++ {
			newPoint := point{
				x: currentPoint.x,
				y: currentPoint.y,
			}
			switch instruction[1] {
			case UP:
				newPoint.y += 1
			case DOWN:
				newPoint.y -= 1
			case RIGHT:
				newPoint.x += 1
			case LEFT:
				newPoint.x -= 1
			}
			if wire1Points[newPoint] {
				crossPoints = append(crossPoints, newPoint)
			}
			currentPoint = newPoint
		}
	}
	return crossPoints
}

func calClosest(points []point) (distance int) {
	closestDistance := math.MaxInt64
	for _, p := range(points) {
		currDist := abs(p.x) + abs(p.y)
		if currDist < closestDistance {
			closestDistance = currDist
		}
	}
	return closestDistance
}

// via https://yourbasic.org/golang/absolute-value-int-float/
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}