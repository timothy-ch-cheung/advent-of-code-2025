package day1

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Instruction struct {
	Direction Direction
	rotations int
}

func Solve() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := parse(scanner)
	solution := solve(instructions)
	fmt.Println(solution)
}

func parse(scanner *bufio.Scanner) []Instruction {
	var instructions []Instruction
	return instructions
}

func solve(instructions []Instruction) string {
	return ""
}
