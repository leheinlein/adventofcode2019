package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type object struct {
	name   string
	orbits *object
}

func main() {
	file, err := os.Open("/Users/eheinlein/go/src/adventOfCode2019/20191206/day6-input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	graph := make(map[string]*object)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		info := strings.Split(txt, ")")
		center := lazyGet(info[0], graph)
		orbiter := lazyGet(info[1], graph)
		orbiter.orbits = center
	}
	var count int
	for _, obj := range graph {
		count += traverse(obj)
	}
	fmt.Println(count)
}

func lazyGet(name string, graph map[string]*object) *object {
	o, success := graph[name]
	if !success {
		newObj := object{name: name}
		o = &newObj
		graph[name] = o
	}
	return o
}

func traverse(obj *object) int {
	if obj.orbits == nil {
		return 0
	}
	return 1 + traverse(obj.orbits)
}
