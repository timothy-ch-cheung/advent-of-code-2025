package day5

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	PATH  = "day5/input.txt"
	TITLE = "Cafeteria"
)

func SolvePart1(path string) int {
	freshRanges, ingredients := parseFreshRanges(path)
	numFresh := 0
	for i := 0; i < len(ingredients); i++ {
		if isFreshLinear(ingredients[i], freshRanges) {
			numFresh++
		}
	}
	return numFresh
}

func parseFreshRanges(path string) ([]freshRange, []int) {
	file, err := os.Open(path)
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	freshRanges := make([]freshRange, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Text() == "" {
			break
		}
		parts := strings.Split(line, "-")
		freshRanges = append(freshRanges, freshRange{
			start: toInt(parts[0]),
			end:   toInt(parts[1]),
		})
	}
	sort.SliceStable(freshRanges, func(i, j int) bool {
		if freshRanges[i].start == freshRanges[j].start {
			return freshRanges[i].end < freshRanges[j].end
		}
		return freshRanges[i].start < freshRanges[j].start
	})

	ingredients := make([]int, 0)
	for scanner.Scan() {
		ingredients = append(ingredients, toInt(scanner.Text()))
	}

	return freshRanges, ingredients
}

type freshRange struct {
	start int
	end   int
}

func isFresh(id int, start int, end int, ranges []freshRange) bool {
	if start > end {
		return false
	}

	mid := intDivide(start+end, 2)
	if mid >= len(ranges) || mid < 0 {
		return false
	}

	if ranges[mid].start <= id && ranges[mid].end >= id {
		return true
	} else if ranges[mid].start > id {
		return isFresh(id, start, mid-1, ranges)
	} else {
		return isFresh(id, mid+1, end, ranges)
	}
}

func isFreshLinear(id int, ranges []freshRange) bool {
	for i := 0; i < len(ranges); i++ {
		if ranges[i].start <= id && ranges[i].end >= id {
			return true
		}
	}
	return false
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func intDivide(numerator int, denominator int) int {
	remainder := numerator % denominator
	return (numerator - remainder) / denominator
}
