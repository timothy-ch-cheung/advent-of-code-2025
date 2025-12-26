package day5

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CafeteriaPart1Suite struct {
	suite.Suite
}

func (c *CafeteriaPart1Suite) TestIsFresh() {
	tests := []struct {
		id       int
		expected bool
	}{
		{id: 3, expected: true},
		{id: 15, expected: true},
		{id: 6, expected: true},
		{id: 7, expected: false},
	}

	ranges := []freshRange{
		{start: 3, end: 5},
		{start: 3, end: 6},
		{start: 10, end: 20},
	}

	for _, tt := range tests {
		c.Run(fmt.Sprintf("[id=%d][expected=[%t]]", tt.id, tt.expected), func() {
			actual := isFresh(tt.id, 0, len(ranges), ranges)
			c.Assert().Equal(tt.expected, actual)
		})
	}
}

func (c *CafeteriaPart1Suite) TestSolvePart1ExampleStepped() {
	tests := []struct {
		id       int
		expected bool
	}{
		{id: 1, expected: false},
		{id: 5, expected: true},
		{id: 8, expected: false},
		{id: 11, expected: true},
		{id: 17, expected: true},
		{id: 32, expected: false},
	}

	ranges := []freshRange{
		{start: 3, end: 5},
		{start: 10, end: 14},
		{start: 12, end: 18},
		{start: 16, end: 20},
	}

	for _, tt := range tests {
		c.Run(fmt.Sprintf("[id=%d][expected=[%t]]", tt.id, tt.expected), func() {
			actual := isFresh(tt.id, 0, len(ranges), ranges)
			c.Assert().Equal(tt.expected, actual)
		})
	}
}

func (c *CafeteriaPart1Suite) TestSolvePart1Example() {
	c.Assert().Equal(3, SolvePart1("example_input.txt"))
}

func TestCafeteriaPart1Suite(t *testing.T) {
	suite.Run(t, new(CafeteriaPart1Suite))
}
