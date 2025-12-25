package day4

import (
	"bufio"
	"os"
)

var adjacentFuncs = []func(int, int, []string) uint8{topLeft, top, topRight, left, right, botLeft, bot, botRight}

func SolvePart1(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	department := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		department = append(department, line)
	}
	for y := 0; y < len(department); y++ {
		for x := 0; x < len(department[y]); x++ {
			if department[y][x] == '@' && isAccessible(x, y, department) {
				count++
			}
		}
	}
	return count
}

func isAccessible(x int, y int, department []string) bool {
	adjacentRolls := 0
	for i := 0; i < len(adjacentFuncs); i++ {
		if adjacentFuncs[i](x, y, department) == '@' {
			adjacentRolls++
		}
	}
	return adjacentRolls < 4
}

func topLeft(x int, y int, department []string) uint8 {
	if x > 0 && y > 0 {
		return department[y-1][x-1]
	}
	return '.'
}

func top(x int, y int, department []string) uint8 {
	if y > 0 {
		return department[y-1][x]
	}
	return '.'
}

func topRight(x int, y int, department []string) uint8 {
	if x < len(department[0])-1 && y > 0 {
		return department[y-1][x+1]
	}
	return '.'
}

func left(x int, y int, department []string) uint8 {
	if x > 0 {
		return department[y][x-1]
	}
	return '.'
}

func right(x int, y int, department []string) uint8 {
	if x < len(department[0])-1 {
		return department[y][x+1]
	}
	return '.'
}

func botLeft(x int, y int, department []string) uint8 {
	if x > 0 && y < len(department)-1 {
		return department[y+1][x-1]
	}
	return '.'
}

func bot(x int, y int, department []string) uint8 {
	if y < len(department)-1 {
		return department[y+1][x]
	}
	return '.'
}

func botRight(x int, y int, department []string) uint8 {
	if x < len(department[0])-1 && y < len(department)-1 {
		return department[y+1][x+1]
	}
	return '.'
}
