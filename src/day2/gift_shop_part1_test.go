package day2

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type GiftShopPart1Suite struct {
	suite.Suite
}

func (s *GiftShopPart1Suite) TestLowerBoundMinLength() {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "1", expected: 1},
		{input: "10", expected: 1},
		{input: "100", expected: 2},
		{input: "1000", expected: 2},
		{input: "10000", expected: 3},
		{input: "100000", expected: 3},
	}

	for _, tt := range tests {
		s.Run("", func() {
			actual := findLowerBoundMinLength(tt.input)
			s.Assert().Equal(tt.expected, actual)
		})
	}
}

func (s *GiftShopPart1Suite) TestLowerBound() {
	tests := []struct {
		input    string
		expected int64
	}{
		{input: "11", expected: 1},
		{input: "95", expected: 9},
		{input: "998", expected: 10},
		{input: "1188511880", expected: 11885},
		{input: "222220", expected: 222},
		{input: "446443", expected: 446},
		{input: "1698522", expected: 1000},
	}

	for _, tt := range tests {
		s.Run("", func() {
			actual := findLowerBound(tt.input)
			s.Assert().Equal(tt.expected, actual)
		})
	}
}

func (s *GiftShopPart1Suite) TestUpperBoundMaxLength() {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "1", expected: 0},
		{input: "10", expected: 1},
		{input: "100", expected: 1},
		{input: "1000", expected: 2},
		{input: "10000", expected: 2},
		{input: "100000", expected: 3},
	}

	for _, tt := range tests {
		s.Run("", func() {
			actual := findUpperBoundMaxLength(tt.input)
			s.Assert().Equal(tt.expected, actual)
		})
	}
}

func (s *GiftShopPart1Suite) TestUpperBound() {
	tests := []struct {
		input    string
		expected int64
	}{
		{input: "22", expected: 2},
		{input: "115", expected: 9},
		{input: "1012", expected: 10},
		{input: "1188511890", expected: 11885},
		{input: "222224", expected: 222},
		{input: "446449", expected: 446},
	}

	for _, tt := range tests {
		s.Run("", func() {
			actual := findUpperBound(tt.input)
			s.Assert().Equal(tt.expected, actual)
		})
	}
}

func (s *GiftShopPart1Suite) TestSumInvalidIds() {
	tests := []struct {
		start    string
		end      string
		expected int64
	}{
		{start: "11", end: "22", expected: 33},
		{start: "95", end: "115", expected: 99},
		{start: "998", end: "1012", expected: 1010},
		{start: "1188511880", end: "1188511890", expected: 1188511885},
		{start: "222220", end: "222224", expected: 222222},
		{start: "1698522", end: "1698528", expected: 0},
	}
	for _, tt := range tests {
		s.Run("", func() {
			actual := sumInvalidIds(tt.start, tt.end)
			s.Assert().Equal(tt.expected, actual)
		})
	}
}

func (s *GiftShopPart1Suite) TestSolvePart1Example() {
	s.Assert().Equal(int64(1227775554), SolvePart1("example_input.txt"))
}

func TestGiftShopPart1Suite(t *testing.T) {
	suite.Run(t, new(GiftShopPart1Suite))
}
