package day1

import (
	"bufio"
	"os"
	"strconv"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type instruction struct {
	direction Direction
	rotations int
}

func SolvePart1(path string) int {
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
			currentPos = rotateLeft(currentPos, currInstruction.rotations)
		case Right:
			currentPos = rotateRight(currentPos, currInstruction.rotations)
		}
		if currentPos == 0 {
			solution += 1
		}
	}

	return solution
}

func rotateLeft(currPos int, rotation int) int {
	rotation = rotation % 100
	diff := currPos - rotation
	if diff < 0 {
		currPos = 100 + diff
	} else {
		currPos = diff
	}
	return currPos
}

func rotateRight(currPos int, rotation int) int {
	rotation = rotation % 100
	sum := currPos + rotation
	if sum < 99 {
		currPos = sum
	} else {
		currPos = sum - 100
	}
	return currPos
}

func parseLine(line string) instruction {
	var direction Direction
	if line[0] == 'R' {
		direction = Right
	} else if line[0] == 'L' {
		direction = Left
	} else {
		panic("Invalid direction code")
	}

	rotations, err := strconv.Atoi(line[1:])
	if err != nil {
		panic("Invalid rotation code")
	}

	return instruction{
		direction: direction,
		rotations: rotations,
	}
}
