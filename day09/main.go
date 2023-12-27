package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1("day09/data.txt")
}

func part1(filename string) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	res := 0
	for _, line := range lines {
		ext := extrapolateBack2(strings.Fields(line))
		fmt.Println(ext)
		res += ext
	}
	fmt.Println(res)
	// fmt.Println(extrapolateBack(lines[0]))
}

func extrapolate(values []string) int {
	stack := make([][]int, 1)
	stack[0] = make([]int, len(values))
	for i, val := range values {
		parsed, _ := strconv.Atoi(val)
		stack[0][i] = parsed
	}

	zeros := false
	last := 0

	for !zeros {
		zeros = true
		next := make([]int, len(stack[last])-1)
		for i := 0; i < len(next); i++ {
			next[i] = stack[last][i+1] - stack[last][i]
			zeros = zeros && (next[i] == 0)
		}
		stack = append(stack, next)
		last++
	}
	// fmt.Printf("%+v\n", stack)
	res := 0
	for _, step := range stack {
		res += step[len(step)-1]
	}

	return res
}

// Thx reddit, again
func extrapolateBack2(fields []string) int {
	slices.Reverse(fields)
	return extrapolate(fields)
}

// This wasn't correct sadly
func extrapolateBack(line string) int {
	values := strings.Fields(line)

	stack := make([][]int, 1)
	stack[0] = make([]int, len(values))
	for i, val := range values {
		parsed, _ := strconv.Atoi(val)
		stack[0][i] = parsed
	}

	zeros := false
	last := 0

	for !zeros {
		zeros = true
		next := make([]int, len(stack[last])-1)
		for i := 0; i < len(next); i++ {
			next[i] = stack[last][i+1] - stack[last][i]
			zeros = zeros && (next[i] == 0)
		}
		stack = append(stack, next)
		last++
	}
	fmt.Printf("%+v\n", stack)
	res := 0
	for _, step := range stack {
		res = step[0] - res
	}

	return res
}
