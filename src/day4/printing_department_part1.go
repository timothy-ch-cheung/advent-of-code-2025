package day4

import (
	"bufio"
	"os"
)

const (
	PATH  = "day4/input.txt"
	TITLE = "Printing Department"
)

const (
	roll = '@'
	free = '.'
)

var adjacentFuncs = []func(int, int, [][]uint8) uint8{topLeft, top, topRight, left, right, botLeft, bot, botRight}

func SolvePart1(path string) int {
	count := 0
	department := parseInput(path)
	for y := 0; y < len(department); y++ {
		for x := 0; x < len(department[y]); x++ {
			if department[y][x] == roll && isAccessible(x, y, department) {
				count++
			}
		}
	}
	return count
}

func parseInput(path string) [][]uint8 {
	file, err := os.Open(path)
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	department := make([][]uint8, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]uint8, len(line))
		department = append(department, row)
		for i := 0; i < len(line); i++ {
			row[i] = line[i]
		}
	}
	return department
}

func isAccessible(x int, y int, department [][]uint8) bool {
	adjacentRolls := 0
	for i := 0; i < len(adjacentFuncs); i++ {
		if adjacentFuncs[i](x, y, department) == roll {
			adjacentRolls++
		}
	}
	return adjacentRolls < 4
}

func topLeft(x int, y int, department [][]uint8) uint8 {
	if x > 0 && y > 0 {
		return department[y-1][x-1]
	}
	return free
}

func top(x int, y int, department [][]uint8) uint8 {
	if y > 0 {
		return department[y-1][x]
	}
	return free
}

func topRight(x int, y int, department [][]uint8) uint8 {
	if x < len(department[0])-1 && y > 0 {
		return department[y-1][x+1]
	}
	return free
}

func left(x int, y int, department [][]uint8) uint8 {
	if x > 0 {
		return department[y][x-1]
	}
	return free
}

func right(x int, y int, department [][]uint8) uint8 {
	if x < len(department[0])-1 {
		return department[y][x+1]
	}
	return free
}

func botLeft(x int, y int, department [][]uint8) uint8 {
	if x > 0 && y < len(department)-1 {
		return department[y+1][x-1]
	}
	return free
}

func bot(x int, y int, department [][]uint8) uint8 {
	if y < len(department)-1 {
		return department[y+1][x]
	}
	return free
}

func botRight(x int, y int, department [][]uint8) uint8 {
	if x < len(department[0])-1 && y < len(department)-1 {
		return department[y+1][x+1]
	}
	return free
}
