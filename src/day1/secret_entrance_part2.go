package day1

import (
	"bufio"
	"os"
)

type rotationResult struct {
	clicks int
	newPos int
}

func SolvePart2(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	solution := 0
	currentPos := 50

	for scanner.Scan() {
		currInstruction := parseLine(scanner.Text())
		switch currInstruction.direction {
		case Left:
			result := rotateLeftV2(currentPos, currInstruction.rotations)
			currentPos = result.newPos
			solution += result.clicks
		case Right:
			result := rotateRightV2(currentPos, currInstruction.rotations)
			currentPos = result.newPos
			solution += result.clicks
		}

		if currentPos == 0 {
			solution += 1
		}
	}

	return solution
}

func rotateLeftV2(currPos int, rotation int) rotationResult {
	var newPos int
	remainder := rotation % 100
	extraClicks := (rotation - remainder) / 100
	diff := currPos - remainder
	if diff < 0 {
		newPos = 100 + diff
		if currPos != 0 && newPos != 0 {
			extraClicks++
		}
	} else {
		newPos = diff
	}
	return rotationResult{
		clicks: extraClicks,
		newPos: newPos,
	}
}

func rotateRightV2(currPos int, rotation int) rotationResult {
	var newPos int
	remainder := rotation % 100
	extraClicks := (rotation - remainder) / 100
	sum := currPos + remainder
	if sum <= 99 {
		newPos = sum
	} else {
		newPos = sum - 100
		if currPos != 0 && newPos != 0 {
			extraClicks++
		}
	}
	return rotationResult{
		clicks: extraClicks,
		newPos: newPos,
	}
}
