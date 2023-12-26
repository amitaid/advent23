package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	part2("day08/data.txt")
}

func part1(filename string) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	directions := lines[0]
	mapLines := lines[2:]
	m := map[string]Step{}

	for _, line := range mapLines {
		// fmt.Printf("%s => %+v\n", line, parseMapLine(line))
		parsed := parseMapLine(line)
		m[parsed.Name] = parsed
	}

	step := m["AAA"]
	count := 0

	for step.Name != "ZZZ" {
		dir := string(directions[count%len(directions)])
		nextStep := step.Next[dir]
		step = m[nextStep]
		count++
	}

	fmt.Println(count, step)

}

type Step struct {
	Name string
	Next map[string]string
}

func parseMapLine(line string) Step {
	name := line[:3]
	left := line[7:10]
	right := line[12:15]
	next := make(map[string]string, 2)
	next["R"] = right
	next["L"] = left
	return Step{Name: name, Next: next}
}

func part2(filename string) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	directions := lines[0]
	mapLines := lines[2:]
	m := map[string]Step{}

	for _, line := range mapLines {
		// fmt.Printf("%s => %+v\n", line, parseMapLine(line))
		parsed := parseMapLine(line)
		m[parsed.Name] = parsed
	}

	paths := []Step{}

	for name, step := range m {
		if strings.HasSuffix(name, "A") {
			paths = append(paths, step)
		}
	}
	fmt.Printf("%s, %+v\n", directions, paths)

	steps := make([]int, len(paths))

	for i, path := range paths {
		count := 0
		for !strings.HasSuffix(path.Name, "Z") {
			dir := string(directions[count%len(directions)])
			path = m[path.Next[dir]]
			count++
		}

		steps[i] = count
	}

	fmt.Println(steps)

	// Thanks, Reddit
	d := 1
	for _, elem := range steps {
		gcd := GCDRemainder(d, elem)
		d *= elem
		d /= gcd
	}

	fmt.Println(d)
}

func GCDRemainder(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
