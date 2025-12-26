package day6

import (
	"advent-of-code-2025/util"
	"bufio"
	"os"
	"strings"
)

const (
	PATH  = "day6/input.txt"
	TITLE = "Trash Compactor"
)

const (
	MULTIPLY = "*"
)

func SolvePart1(path string) int {
	file, err := os.Open(path)
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	worksheet := make([][]string, 0)
	for scanner.Scan() {
		worksheet = append(worksheet, strings.Fields(scanner.Text()))
	}
	return solveWorksheet(worksheet)
}

func solveWorksheet(worksheet [][]string) int {
	operatorIdx := len(worksheet) - 1
	sum := 0

	for x := 0; x < len(worksheet[0]); x++ {
		operator := worksheet[operatorIdx][x]
		result := 0
		if operator == MULTIPLY {
			result = 1
		}
		for y := 0; y < len(worksheet)-1; y++ {
			if operator == MULTIPLY {
				result *= util.ToInt(worksheet[y][x])
			} else {
				result += util.ToInt(worksheet[y][x])
			}
		}
		sum += result
	}
	return sum
}
