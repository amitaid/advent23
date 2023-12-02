package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	part2("day01/data.txt")
}

func part1(filename string) {
	data, _ := os.ReadFile(filename)
	var res, first, last int
	seenFirst := false
	for _, b := range data {
		if n, ok := getDigit(b); ok {
			if !seenFirst {
				first = n
				seenFirst = true
			}
			last = n
		} else if b == '\n' && seenFirst {
			res += 10*first + last
			seenFirst = false
		}
	}
	println(res)
}

func getDigit(b byte) (int, bool) {
	num := int(b - '0')
	return num, num >= 0 && num <= 9
}

var numbers = map[string]int{"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func hasNumber(sub string) (int, bool) {
	l := len(sub)
	if l == 0 {
		return -1, false
	}

	if n, ok := getDigit(sub[0]); ok {
		return n, ok
	}

	if l < 3 {
		return -1, false
	}

	for i := 5; i >= 3; i-- {
		if l >= i {
			if num, ok := numbers[sub[:i]]; ok {
				return num, ok
			}
		}
	}
	return -1, false
}

func part2(filename string) {
	data, _ := os.ReadFile(filename)
	parts := bytes.Split(data, []byte("\n"))
	var res int

	for _, part := range parts {
		line := string(part)
		l := len(line)
		fmt.Println(line)
		for i := 0; i < l; i++ {
			if n, ok := hasNumber(line[i:]); ok {
				fmt.Printf("first %d\n", n)
				res += 10 * n
				break
			}
		}
		for i := l - 1; i >= 0; i-- {
			if n, ok := hasNumber(line[i:]); ok {
				fmt.Printf("last %d\n", n)
				res += n
				break
			}
		}
	}
	println(res)
}
