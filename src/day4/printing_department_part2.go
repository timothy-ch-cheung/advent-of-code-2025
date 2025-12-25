package day4

const removed = 'X'

func SolvePart2(path string) int {
	department := parseInput(path)
	totalRemoved := 0
	removedRolls := removeRolls(department)
	for removedRolls != 0 {
		totalRemoved += removedRolls
		removedRolls = removeRolls(department)
	}
	return totalRemoved
}

func removeRolls(department [][]uint8) int {
	removedRolls := 0
	for y := 0; y < len(department); y++ {
		for x := 0; x < len(department[y]); x++ {
			if department[y][x] == roll && isAccessible(x, y, department) {
				department[y][x] = removed
				removedRolls++
			}
		}
	}
	return removedRolls
}
