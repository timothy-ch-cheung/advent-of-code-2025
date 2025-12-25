package day4

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PrintingDepartmentPart1Suite struct {
	suite.Suite
}

func (p *PrintingDepartmentPart1Suite) TestPrintingDepartmentPart1() {
	tests := []struct {
		directionFunc func(int, int, [][]uint8) uint8
		x             int
		y             int
		expected      uint8
	}{
		{directionFunc: topLeft, x: 1, y: 1, expected: '1'},
		{directionFunc: topLeft, x: 0, y: 0, expected: '.'},
		{directionFunc: top, x: 1, y: 1, expected: '2'},
		{directionFunc: top, x: 1, y: 0, expected: '.'},
		{directionFunc: topRight, x: 1, y: 1, expected: '3'},
		{directionFunc: topRight, x: 2, y: 0, expected: '.'},
		{directionFunc: left, x: 1, y: 1, expected: '4'},
		{directionFunc: left, x: 0, y: 1, expected: '.'},
		{directionFunc: right, x: 1, y: 1, expected: '6'},
		{directionFunc: right, x: 2, y: 1, expected: '.'},
		{directionFunc: botLeft, x: 1, y: 1, expected: '7'},
		{directionFunc: botLeft, x: 0, y: 2, expected: '.'},
		{directionFunc: bot, x: 1, y: 1, expected: '8'},
		{directionFunc: bot, x: 1, y: 2, expected: '.'},
		{directionFunc: botRight, x: 1, y: 1, expected: '9'},
		{directionFunc: botRight, x: 2, y: 2, expected: '.'},
	}
	department := [][]uint8{
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', '9'},
	}

	for _, tt := range tests {
		p.Run(fmt.Sprintf("[x=%d][y=%d][%s]", tt.x, tt.y, string(tt.expected)), func() {
			actual := tt.directionFunc(tt.x, tt.y, department)
			p.Assert().Equal(tt.expected, actual)
		})
	}
}

func (p *PrintingDepartmentPart1Suite) TestSolvePart1Example() {
	p.Assert().Equal(13, SolvePart1("example_input.txt"))
}

func TestPrintingDepartmentPart1Suite(t *testing.T) {
	suite.Run(t, new(PrintingDepartmentPart1Suite))
}
