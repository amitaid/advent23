package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part2("day05/data.txt")
}

type Map struct {
	From, To string
	Mappings []Mapping
}

type Mapping struct {
	Start, End, Offset int
}

func part1(filename string) {
	data, _ := os.ReadFile(filename)
	segments := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	maps := parseMaps(segments[1:])

	seedsStr := strings.Split(segments[0][7:], " ")

	lowest := 9223372036854775807
	for _, seedStr := range seedsStr {
		seed, _ := strconv.Atoi(seedStr)
		location := processSeed(seed, maps)
		if location < lowest {
			lowest = location
		}
	}

	fmt.Println(lowest)

}

func processSeed(seed int, maps map[string]Map) int {
	mapType := "seed"
	res := seed

	fmt.Printf("%s %d", mapType, res)
	for mapType != "location" {
		m := maps[mapType]
		res = m.Apply(res)
		mapType = m.To
		fmt.Printf(" -> %s %v %d\n", mapType, m, res)
	}
	fmt.Println()
	return res
}

func parseMaps(mapsRaw []string) (maps map[string]Map) {
	maps = make(map[string]Map, len(mapsRaw))

	for _, raw := range mapsRaw {
		m := parseMap(raw)
		maps[m.From] = m
	}

	return
}

func parseMap(raw string) Map {
	mapParts := strings.Split(raw, "\n")
	fromTo := strings.Split(mapParts[0][:strings.Index(mapParts[0], " ")], "-")
	res := Map{From: fromTo[0], To: fromTo[2]}

	for _, rawMap := range mapParts[1:] {
		rawMapParts := strings.Split(rawMap, " ")

		m := Mapping{}
		m.Start, _ = strconv.Atoi(rawMapParts[1])

		m.End, _ = strconv.Atoi(rawMapParts[2])
		m.End += m.Start

		m.Offset, _ = strconv.Atoi(rawMapParts[0])
		m.Offset -= m.Start
		res.Mappings = append(res.Mappings, m)
	}
	slices.SortFunc(res.Mappings, func(a, b Mapping) int { return cmp.Compare(a.Start, b.Start) })

	return res
}

func (m Map) Apply(element int) int {
	// fmt.Printf("applying %v on %d\n", m, element)
	for _, mapping := range m.Mappings {
		if element < mapping.Start {
			return element
		} else if element < mapping.End {
			return element + mapping.Offset
		}
	}
	return element
}

type Range struct {
	Start, End int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (r Range) Merge(other Range) (Range, bool) {
	if r.End < other.Start {
		return r, false
	}
	return Range{r.Start, max(r.End, other.End)}, true
}

func mergeRanges(ranges []Range) []Range {
	// fmt.Printf("merging %v\n", ranges)
	slices.SortFunc(ranges, func(a, b Range) int { return cmp.Compare[int](a.Start, b.Start) })

	res := []Range{}
	i := 0
	for i < len(ranges) {
		cur := ranges[i]
		merged := true
		j := i + 1
		for j < len(ranges) && merged {
			if cur, merged = cur.Merge(ranges[j]); merged {
				j += 1
			}
		}

		res = append(res, cur)
		i = j
	}

	// fmt.Printf("result %v\n", res)
	return res
}

func part2(filename string) {
	data, _ := os.ReadFile(filename)
	segments := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	maps := parseMaps(segments[1:])

	seedRanges := parseSeedRanges(segments[0])

	fmt.Println(seedRanges)

	// fmt.Println(maps["seed"].ApplyRanges(seedRanges))

	targetRanges := processSeedRanges(maps, seedRanges)

	lowest := 9223372036854775807
	for _, r := range targetRanges {
		if r.Start < lowest {
			lowest = r.Start
		}
	}
	fmt.Println(lowest)
}

func processSeedRanges(maps map[string]Map, ranges []Range) []Range {
	step := "seed"
	applied := []Range{}
	unapplied := ranges

	for step != "location" {
		m := maps[step]
		for _, mm := range m.Mappings {
			fmt.Printf("applying %v to %v\n", mm, unapplied)
			var temp []Range
			for _, r := range unapplied {
				mapped, unmapped := mm.IntersectRange(r)
				// fmt.Printf("%v x %v => %v, %v\n", mm, r, mapped, unmapped)
				applied = append(applied, mapped...)
				temp = append(temp, unmapped...)
			}
			unapplied = temp

			// fmt.Printf("applied %v unapplied %v\n", applied, unapplied)
		}
		step = m.To
		unapplied = mergeRanges(append(applied, unapplied...))
		applied = []Range{}
		fmt.Printf("============= step %s result %v\n", m.From, unapplied)
	}

	// fmt.Printf("result %v\n", unapplied)
	return unapplied
}

// func processMap(maps map[string]Map, r Range, step string) []Range {
// 	if step == "location" {
// 		return []Range{r}
// 	}
// 	// fmt.Printf("Applying map %s: %v\n", step, maps[step])

// 	m := maps[step]
// 	for _, mm := range m.Mappings {
// 		var temp []Range
// 		for _, rr := range unapplied {
// 			mapped, unmapped := mm.IntersectRange(r)
// 			applied = append(applied, mapped...)
// 			temp = append(temp, unmapped...)
// 		}
// 		unapplied = temp

// 	}
// 	if len(applied) == 0 {
// 		applied = []Range{r}
// 		fmt.Printf("%v x %v => [%v]\n", m, r, r)
// 	}

// 	res := []Range{}
// 	for _, rr := range applied {
// 		res = append(res, processRange(maps, rr, m.To)...)
// 	}

// 	return mergeRanges(res)
// }

func parseSeedRanges(seedLine string) []Range {
	seedLine = strings.TrimSpace(seedLine[7:])
	seedRanges := []Range{}
	for len(seedLine) > 0 {
		split := strings.SplitN(seedLine, " ", 3)
		start, _ := strconv.Atoi(split[0])
		length, _ := strconv.Atoi(split[1])
		if len(split) == 2 {
			seedLine = ""
		} else {
			seedLine = split[2]
		}

		seedRanges = append(seedRanges, Range{Start: start, End: start + length})
	}
	return seedRanges
}

func (m Mapping) IntersectRange(r Range) (mapped, unmapped []Range) {
	if r.End <= m.Start || r.Start >= m.End {
		// fmt.Printf("range %v was outside %v\n", r, m)
		return []Range{}, []Range{r}
	}

	bot := r.Start
	top := r.End

	if r.Start < m.Start {
		// fmt.Printf("range %v was below %v ", r, m)
		unmapped = append(unmapped, Range{r.Start, m.Start})
		bot = m.Start
	}
	if r.End > m.End {
		// fmt.Printf("range %v was above %v ", r, m)
		unmapped = append(unmapped, Range{m.End, r.End})
		top = m.End
	}

	// fmt.Printf("range %v intersected with %v ", r, m)
	mapped = append(mapped, Range{bot + m.Offset, top + m.Offset})

	// fmt.Printf("%v x %v => %v\n", m, r, res)
	return
}
