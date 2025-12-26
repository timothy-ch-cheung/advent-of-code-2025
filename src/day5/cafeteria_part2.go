package day5

import (
	"advent-of-code-2025/util"
)

func SolvePart2(path string) int {
	freshRanges, _ := parseFreshRanges(path)

	freshRanges = mergeRanges(freshRanges)

	numFresh := 0
	for i := 0; i < len(freshRanges); i++ {
		numFresh += freshRanges[i].end - freshRanges[i].start + 1
	}

	return numFresh
}

func mergeRanges(freshRanges []freshRange) []freshRange {
	hasMerged := true
	for hasMerged {
		hasMerged = false
		newFreshRanges := make([]freshRange, 0)
		for i := 0; i < len(freshRanges); i++ {
			if i < len(freshRanges)-1 && freshRanges[i].end >= freshRanges[i+1].start {
				newFreshRanges = append(newFreshRanges, freshRange{
					start: freshRanges[i].start,
					end:   util.Max(freshRanges[i].end, freshRanges[i+1].end),
				})
				hasMerged = true
				i++
			} else {
				newFreshRanges = append(newFreshRanges, freshRanges[i])
			}
		}
		freshRanges = newFreshRanges
	}
	return freshRanges
}
