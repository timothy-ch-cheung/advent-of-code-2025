package day2

import (
	"bufio"
	"fmt"
	"os"
)

func SolvePart2(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	productRanges := parseLine(scanner.Text())

	var sum int64 = 0
	for i := 0; i < len(productRanges); i++ {
		sum += sumInvalidProductIdsForAllLengths(productRanges[i].start, productRanges[i].end)
	}
	return sum
}

func sumInvalidProductIdsForAllLengths(start string, end string) int64 {
	seenMap := make(map[int64]bool)
	lengths := intDivide(len(end), 2)
	var total int64 = 0
	for i := 1; i <= lengths; i++ {
		lowerBound := findLowerBoundGivenLength(start, i)
		upperBound := findUpperBoundGivenLength(end, i)
		fmt.Printf("start[%s] end[%s] lower[%d] upper[%d] length[%d] repeats[%d]\n", start, end, lowerBound, upperBound, lengths, intDivide(len(end), lengths))
		for j := lowerBound; j <= upperBound; j++ {
			part := toStr(j)
			repeats := intDivide(len(end), lengths)
			repeated := toInt(repeat(part, repeats))
			if _, ok := seenMap[repeated]; !ok {
				total += repeated
				seenMap[repeated] = true
			}
		}
	}
	return total
}

func findLowerBoundGivenLength(start string, length int) int64 {
	repeatDigits := start[:length]
	repeats := intDivide(len(start), length)
	if length*repeats < len(start) {
		repeats++
	}
	lowerBound := toInt(repeat(repeatDigits, repeats))
	if lowerBound < toInt(start) && toInt(repeatDigits) < int64(pow(10, len(repeatDigits))-1) {
		return toInt(toStr(toInt(repeatDigits) + 1))
	}
	return toInt(repeatDigits)
}

func findUpperBoundGivenLength(end string, length int) int64 {
	if len(end)%length != 0 {
		return toInt(repeat("9", length))
	}
	repeatDigits := end[:length]
	repeats := intDivide(len(end), length)
	upperBound := toInt(repeat(repeatDigits, repeats))
	if upperBound > toInt(end) {
		return toInt(repeatDigits) - 1
	}
	return toInt(repeatDigits)
}
