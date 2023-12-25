package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part2("day07/data.txt")
}

func part1(filename string) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	bids := parseBids(lines, false)

	sort.Slice(bids, func(i, j int) bool { return compare(bids[i].hand, bids[j].hand) < 0 })

	res := 0
	for i, bid := range bids {
		res += (i + 1) * bid.amount
		fmt.Printf("%+v\n", bid)
	}
	fmt.Println(res)
}

var cards = map[byte]int{
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func cardStrength(card byte) int {
	str, ok := cards[card]
	if ok {
		return str
	}
	return int(card - '0')
}

type Category struct {
	Name string
	Rank int
}

type Hand struct {
	cards    string
	category Category
	counts   [][]int
}

func parseHand(hand string, jokersEnabled bool) Hand {
	res := make([][]int, 5)
	cardCounts := map[int]int{}
	jokers := 0
	for i := 0; i < 5; i++ {
		if jokersEnabled && hand[i] == 'J' {
			jokers++
		} else {
			str := cardStrength(hand[i])
			if _, ok := cardCounts[str]; !ok {
				cardCounts[str] = 0
			}
			cardCounts[str]++
		}
	}

	for card, count := range cardCounts {
		res[5-count] = append(res[5-count], card)
	}
	for _, row := range res {
		sort.Slice(row, func(i, j int) bool { return row[i] > row[j] })
	}

	return Hand{cards: hand, counts: res, category: getCategory(res, jokers)}
}

func getCategory(counts [][]int, jokers int) Category {
	if len(counts[4]) == 5 { // High card
		return Category{Name: "High Card", Rank: 0}
	}
	if len(counts[3]) == 1 && len(counts[2]) == 0 { // One pair
		if jokers == 1 {
			return Category{Name: "Three", Rank: 30} // Three
		} else if jokers == 2 {
			return Category{Name: "Four", Rank: 40} // Four
		} else if jokers == 3 {
			return Category{Name: "Five", Rank: 50} // Five
		}
		return Category{Name: "Pair", Rank: 20}
	}
	if len(counts[3]) == 2 { // Two pairs
		if jokers == 1 {
			return Category{Name: "Full", Rank: 31} // Full house
		}
		return Category{Name: "Two Pairs", Rank: 21}
	}
	if len(counts[2]) == 1 && len(counts[3]) == 1 { // Full house, can't have jokers
		return Category{Name: "Full", Rank: 31}
	}
	if len(counts[2]) == 1 { // Three
		if jokers == 1 {
			return Category{Name: "Four", Rank: 40} // Four
		} else if jokers == 2 {
			return Category{Name: "Five", Rank: 50} // Five
		}
		return Category{Name: "Three", Rank: 30}
	}
	if len(counts[1]) == 1 { // Four
		if jokers == 1 {
			return Category{Name: "Five", Rank: 50}
		}
		return Category{Name: "Four", Rank: 40}
	}
	if len(counts[0]) == 1 { // Five
		return Category{Name: "Five", Rank: 50}
	}
	if len(counts[4]) > 0 {
		return Category{Name: "High Card", Rank: 10 * (jokers + 1)}
	}
	return Category{Name: "Jokers", Rank: 10 * jokers}
}

func compare(h1, h2 Hand) int {
	if h1.category.Rank != h2.category.Rank {
		return h1.category.Rank - h2.category.Rank
	}
	cmp := 0
	for i := 0; i < 5 && cmp == 0; i++ {
		cmp = cardStrength(h1.cards[i]) - cardStrength(h2.cards[i])
	}
	return cmp
}

type Bid struct {
	hand   Hand
	amount int
}

func parseBids(lines []string, jokers bool) []Bid {
	res := make([]Bid, 0, len(lines))
	for _, line := range lines {
		parts := strings.Fields(line)
		n, _ := strconv.Atoi(parts[1])
		bid := Bid{hand: parseHand(parts[0], jokers), amount: n}
		res = append(res, bid)
	}
	return res
}

func part2(filename string) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	cards['J'] = 1 // Jokers have lowest prio

	bids := parseBids(lines, true)
	sort.Slice(bids, func(i, j int) bool { return compare(bids[i].hand, bids[j].hand) < 0 })

	res := 0
	for i, bid := range bids {
		res += (i + 1) * bid.amount
		fmt.Printf("%+v\n", bid)
	}
	fmt.Println(res)
}
