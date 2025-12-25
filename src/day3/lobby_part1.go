package day3

import (
	"bufio"
	"os"
	"strconv"
)

const INT_OFFSET = 48

func SolvePart1(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		bank := scanner.Text()
		total += maxJoltage(bank)
	}
	return total
}

func maxJoltage(bank string) int {
	firstMax := 0
	firstMaxIdx := 0
	for i := 0; i < len(bank)-1; i++ {
		current := int(bank[i] - INT_OFFSET)
		if current > firstMax {
			firstMax = current
			firstMaxIdx = i
		}
	}
	secondMax := 0
	for i := firstMaxIdx + 1; i < len(bank); i++ {
		current := int(bank[i] - INT_OFFSET)
		if current > secondMax {
			secondMax = current
		}
	}
	return firstMax*10 + secondMax
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}
