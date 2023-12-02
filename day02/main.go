package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2("day02/data.txt")
}

type Round struct {
	cubes map[string]int
}

type Game struct {
	id     int
	rounds []Round
}

func part1(filename string) {
	f, _ := os.Open(filename)
	data := bufio.NewScanner(f)
	var res int

	for data.Scan() {
		game := parseGame(data.Text())
		if game.isPossible() {
			res += game.id
		}
	}
	fmt.Println(res)
}

func part2(filename string) {
	f, _ := os.Open(filename)
	data := bufio.NewScanner(f)
	var res int

	for data.Scan() {
		game := parseGame(data.Text())
		res += game.power()
	}
	fmt.Println(res)
}

var limits = map[string]int{"red": 12, "green": 13, "blue": 14}

func (g Game) isPossible() bool {
	for _, round := range g.rounds {
		for color, amount := range limits {
			if num, ok := round.cubes[color]; ok && num > amount {
				return false
			}
		}
	}
	return true
}

func (g Game) findMinCubes() map[string]int {
	res := map[string]int{"red": 0, "blue": 0, "green": 0}
	for _, round := range g.rounds {
		for color, amount := range round.cubes {
			if res[color] < amount {
				res[color] = amount
			}
		}
	}
	return res
}

func (g Game) power() int {
	min := g.findMinCubes()
	res := 1
	for _, num := range min {
		res *= num
	}
	return res
}

func parseRounds(s string) []Round {
	rs := strings.Split(s, "; ")
	rounds := make([]Round, 0, len(rs))
	for _, r := range rs {
		cubes := strings.Split(r, ", ")
		round := Round{cubes: map[string]int{"red": 0, "blue": 0, "green": 0}}
		for _, c := range cubes {
			nc := strings.Split(c, " ")
			n, _ := strconv.ParseUint(nc[0], 10, 64)
			round.cubes[nc[1]] = int(n)
		}
		rounds = append(rounds, round)
	}
	return rounds
}

func parseGame(s string) Game {
	gameAndRounds := strings.Split(s, ": ")
	id, _ := strconv.ParseUint(gameAndRounds[0][5:], 10, 32)
	return Game{
		id:     int(id),
		rounds: parseRounds(gameAndRounds[1]),
	}
}
