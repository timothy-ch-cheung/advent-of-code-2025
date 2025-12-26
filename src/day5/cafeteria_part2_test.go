package day5

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CafeteriaPart2Suite struct {
	suite.Suite
}

func (c *CafeteriaPart2Suite) TestMergeRanges() {
	tests := []struct {
		input    []freshRange
		expected []freshRange
	}{
		{
			input:    []freshRange{{3, 5}, {10, 14}, {12, 18}, {16, 20}},
			expected: []freshRange{{3, 5}, {10, 20}},
		},
		{
			input:    []freshRange{{3, 5}, {3, 20}},
			expected: []freshRange{{3, 20}},
		},
		{
			input:    []freshRange{{1, 500}, {1, 20}, {2, 40}},
			expected: []freshRange{{1, 500}},
		},
	}

	for _, tt := range tests {
		c.Run(fmt.Sprintf("[input=%d][expected=[%d]]", tt.input, tt.expected), func() {
			actual := mergeRanges(tt.input)
			c.Assert().Equal(tt.expected, actual)
		})
	}
}

func (c *CafeteriaPart2Suite) TestSolvePart2Example() {
	c.Assert().Equal(14, SolvePart2("example_input.txt"))
}

func TestCafeteriaPart2Suite(t *testing.T) {
	suite.Run(t, new(CafeteriaPart2Suite))
}
