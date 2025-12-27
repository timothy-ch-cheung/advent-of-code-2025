package day6

import (
	"advent-of-code-2025/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ADD_CHAR  = '+'
	MULT_CHAR = '*'
)

func SolvePart2(path string) int {
	worksheet := parseWorksheetV2(path)
	return solveWorksheetV2(worksheet)
}

func parseWorksheetV2(path string) []string {
	file, err := os.Open(path)
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	worksheet := make([]string, 0)
	for scanner.Scan() {
		worksheet = append(worksheet, scanner.Text())
	}
	width := maxLength(worksheet)
	for i := 0; i < len(worksheet); i++ {
		if len(worksheet[i]) < width {
			worksheet[i] = fmt.Sprintf("%-"+strconv.Itoa(width)+"s", worksheet[i])
		}
	}
	return worksheet
}

func solveWorksheetV2(worksheet []string) int {
	sum := 0

	numbers := make([]int, 0)
	var operator uint8
	for x := 0; x < len(worksheet[0]); x++ {
		column := ""
		for y := 0; y < len(worksheet); y++ {
			column += string(worksheet[y][x])
		}

		column = strings.TrimSpace(column)
		if len(column) == 0 || x == len(worksheet[0])-1 {
			if len(column) != 0 {
				numbers = append(numbers, util.ToInt(strings.TrimSpace(column)))
			}
			var result int
			if operator == MULT_CHAR {
				result = 1
			} else {
				result = 0
			}
			for i := 0; i < len(numbers); i++ {
				if operator == MULT_CHAR {
					result *= numbers[i]
				} else {
					result += numbers[i]
				}
			}
			sum += result
			numbers = make([]int, 0)
		} else if util.Last(column) == ADD_CHAR || util.Last(column) == MULT_CHAR {
			operator = column[len(column)-1]
			column = column[:len(column)-1]
			column = strings.TrimSpace(column)
			numbers = append(numbers, util.ToInt(column))
		} else {
			numbers = append(numbers, util.ToInt(column))
		}
	}
	return sum
}

func maxLength(worksheet []string) int {
	maxLen := 0
	for i := 0; i < len(worksheet); i++ {
		if len(worksheet[i]) > maxLen {
			maxLen = len(worksheet[i])
		}
	}
	return maxLen
}
