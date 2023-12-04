package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part2("day03/data.txt")
}

func part1(filename string) {
	data, _ := os.ReadFile(filename)
	lines := bytes.Split(data, []byte("\n"))

	var currentNum []byte
	var currentHasSymbol bool
	var res int64
	count := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if (!isDigit(lines[i][j]) || j == 0) && currentNum != nil {
				if currentHasSymbol {
					num, _ := strconv.ParseInt(string(currentNum), 10, 64)
					res += num
					count++
					currentHasSymbol = false
				}
				currentNum = nil
			}
			if isDigit(lines[i][j]) {
				currentNum = append(currentNum, lines[i][j])
				currentHasSymbol = currentHasSymbol || hasNearby(lines, i, j, isSymbol)
			}
		}
	}
	fmt.Printf("res %d", res)
}

func hasNearby(lines [][]byte, row, col int, finder func(byte) bool) bool {
	// fmt.Printf("testing %d,%d\n", row, col)
	if row > 0 {
		if col > 0 && finder(lines[row-1][col-1]) {
			return true
		}
		if finder(lines[row-1][col]) {
			return true
		}
		if col < len(lines[row-1])-1 && finder(lines[row-1][col+1]) {
			return true
		}
	}
	if row < len(lines)-1 {
		if col > 0 && finder(lines[row+1][col-1]) {
			return true
		}
		if finder(lines[row+1][col]) {
			return true
		}
		if col < len(lines[row+1])-1 && finder(lines[row+1][col+1]) {
			return true
		}
	}

	return col > 0 && isSymbol(lines[row][col-1]) || col < len(lines[row])-1 && isSymbol(lines[row][col+1])
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSymbol(b byte) bool {
	// fmt.Printf("testing %c\n", b)
	return b != '.' && !isDigit(b)
}

func part2(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		println(err.Error())
	}
	lines := bytes.Split(data, []byte("\n"))

	var res int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '*' {
				if mul, ok := getNearbyTwoNumsMul(lines, i, j); ok {
					res += mul
				}
			}
		}
	}
	println(res)
}

func getNearbyTwoNumsMul(lines [][]byte, row, col int) (int, bool) {
	var res []int
	if num, ok := tryReadNum(lines, row-1, col); ok {
		res = append(res, num)
	} else {
		if num, ok = tryReadNum(lines, row-1, col-1); ok {
			res = append(res, num)
		}
		if num, ok = tryReadNum(lines, row-1, col+1); ok {
			res = append(res, num)
		}
	}

	if num, ok := tryReadNum(lines, row+1, col); ok {
		res = append(res, num)
	} else {
		if num, ok = tryReadNum(lines, row+1, col-1); ok {
			res = append(res, num)
		}
		if num, ok = tryReadNum(lines, row+1, col+1); ok {
			res = append(res, num)
		}
	}

	if num, ok := tryReadNum(lines, row, col-1); ok {
		res = append(res, num)
	}
	if num, ok := tryReadNum(lines, row, col+1); ok {
		res = append(res, num)
	}

	if len(res) == 2 {
		return res[0] * res[1], true
	}
	return 0, false
}

func tryReadNum(lines [][]byte, i, j int) (int, bool) {
	if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[i]) {
		return 0, false
	}

	l := len(lines[i])
	var res int
	if isDigit(lines[i][j]) {
		res += int(lines[i][j] - '0')
		// read left
		if j-1 >= 0 && isDigit(lines[i][j-1]) {
			res += 10 * int(lines[i][j-1]-'0')
			if j-2 >= 0 && isDigit(lines[i][j-2]) {
				res += 100 * int(lines[i][j-2]-'0')
			}
		}
		// read right
		if j+1 < l && isDigit(lines[i][j+1]) {
			res = 10*res + int(lines[i][j+1]-'0')
			if j+2 < l && isDigit(lines[i][j+2]) {
				res = 10*res + int(lines[i][j+2]-'0')
			}
		}
		return res, true
	}

	return 0, false
}
