package day3

import (
	"bufio"
	"os"
)

func SolvePart2(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int64 = 0
	for scanner.Scan() {
		bank := scanner.Text()
		total += maxJoltageV2(bank, 12)
	}
	return total
}

func maxJoltageV2(bank string, n int) int64 {
	if n == 0 {
		return 0
	}
	var currMax int64 = 0
	var currMaxIdx int64 = 0

	for i := 0; i < len(bank)-n+1; i++ {
		var current = int64(bank[i] - INT_OFFSET)
		if current > currMax {
			currMax = current
			currMaxIdx = int64(i)
		}
	}
	return currMax*pow(10, int64(n-1)) + maxJoltageV2(bank[currMaxIdx+1:], n-1)
}

func pow(base, exp int64) int64 {
	var result int64 = 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}
