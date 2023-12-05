package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	part2("day04/data.txt")
}

func part1(filename string) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	var res int
	for _, card := range lines {
		res += processCard(card)
	}
	println(res)
}

func parseCards(cards []string) []int {
	res := make([]int, len(cards))

	for i, card := range cards {
		colon := strings.Index(card, ":")
		winningAndNums := strings.Split(card[colon+2:], " | ")
		winning := map[string]bool{}
		winningRaw := strings.Split(winningAndNums[0], " ")
		for _, num := range winningRaw {
			winning[num] = len(num) > 0
		}

		nums := strings.Split(winningAndNums[1], " ")
		for _, num := range nums {
			if win, ok := winning[num]; win && ok {
				res[i] += 1
			}
		}
	}
	return res
}

func processCard(card string) int {
	colon := strings.Index(card, ":")
	winningAndNums := strings.Split(card[colon+2:], " | ")
	winning := map[string]bool{}
	winningRaw := strings.Split(winningAndNums[0], " ")
	for _, num := range winningRaw {
		winning[num] = len(num) > 0
	}

	var res int
	nums := strings.Split(winningAndNums[1], " ")
	for _, num := range nums {
		if winning[num] {
			if res == 0 {
				res = 1
			} else {
				res *= 2
			}
		}
	}

	return res
}

func part2(filename string) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	cards := parseCards(lines)

	res := 0
	cardCounts := make([]int, len(lines))
	for i := range cardCounts {
		res += cardCounts[i] + 1
		for j := 1; i+j < len(cards) && j < cards[i]+1; j++ {
			cardCounts[i+j] += (cardCounts[i] + 1)
		}
		// fmt.Println(cardCounts)
	}
	fmt.Println(res)
}
