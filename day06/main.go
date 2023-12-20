package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2("day06/data.txt")
}

func part1(filename string) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	races := parse(lines)

	res := 1
	for _, r := range races {
		race := handleRace(r)
		res *= race
	}

	fmt.Println(res)
}

type Race struct {
	Time, Distance float64
}

func parse(lines []string) []Race {
	timesStr := strings.Fields(lines[0])[1:]
	distStr := strings.Fields(lines[1])[1:]

	res := []Race{}
	for i := range timesStr {
		time, _ := strconv.Atoi(timesStr[i])
		dist, _ := strconv.Atoi(distStr[i])
		res = append(res, Race{float64(time), float64(dist)})
	}

	return res
}

// Race wins are between Time/2 +- sqrt((Time/2)^2 - Distance). Slight adjustment for when the top end is an integer.
func handleRace(r Race) int {
	halfT := r.Time / 2
	min := halfT - math.Sqrt(halfT*halfT-r.Distance)
	max := halfT + math.Sqrt(halfT*halfT-r.Distance)
	res := int(math.Floor(max) - math.Floor(min))

	if max == float64(int(max)) {
		res -= 1
	}

	fmt.Printf("race %v had %d ways to win, min %f max %f\n", r, res, min, max)

	return res

}

func part2(filename string) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	race := parse2(lines)
	fmt.Println(handleRace(race))
}

func parse2(lines []string) Race {
	t, _ := strconv.Atoi(strings.ReplaceAll(lines[0][9:], " ", ""))
	d, _ := strconv.Atoi(strings.ReplaceAll(lines[1][9:], " ", ""))
	return Race{float64(t), float64(d)}
}
