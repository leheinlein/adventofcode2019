package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type object struct {
	name     string
	orbiters map[*object]bool
}

func main() {
	file, err := os.Open("/Users/eheinlein/go/src/adventOfCode2019/20191206/day6-input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	graph := make(map[string]*object)
	var start, target *object

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		info := strings.Split(txt, ")")
		center := lazyGet(info[0], graph)
		orbiter := lazyGet(info[1], graph)
		center.orbiters[orbiter] = true
		orbiter.orbiters[center] = true
		switch orbiter.name {
		case "YOU":
			start = center
		case "SAN":
			target = center
		}
	}
	count := findLengthToTarget(start, target)
	fmt.Println(count)
}

func findLengthToTarget(start *object, target *object) int {
	shortestPaths := make(map[*object]int)
	traverseGraph(start, shortestPaths, 0)
	return shortestPaths[target]

}

func traverseGraph(curr *object, shortestPaths map[*object]int, currDistance int) {
	for obj := range curr.orbiters {
		newDist := currDistance + 1
		if shortestPaths[obj] == 0 || shortestPaths[obj] > newDist {
			shortestPaths[obj] = newDist
			traverseGraph(obj, shortestPaths, newDist)
		}
	}
}

func lazyGet(name string, graph map[string]*object) *object {
	o, success := graph[name]
	if !success {
		newObj := object{
			name:     name,
			orbiters: make(map[*object]bool),
		}
		o = &newObj
		graph[name] = o
	}
	return o
}
