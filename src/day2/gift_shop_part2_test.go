package day2

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GiftShopPart2Suite struct {
	suite.Suite
}

func (s *GiftShopPart2Suite) TestFindLowerBounds() {
	tests := []struct {
		start    string
		length   int
		expected int64
	}{
		{start: "11", length: 1, expected: 1},
		{start: "95", length: 1, expected: 1},
		{start: "998", length: 1, expected: 9},
		{start: "1188511880", length: 5, expected: 11885},
		{start: "222220", length: 3, expected: 222},
		{start: "222220", length: 1, expected: 2},
		{start: "1698522", length: 1, expected: 2},
		{start: "998", length: 2, expected: 10},
	}
	for _, tt := range tests {
		s.Run(fmt.Sprintf("start:%s", tt.start), func() {
			actual := findLowerBoundGivenLength(tt.start, tt.length)
			s.Assert().Equal(tt.expected, actual)
		})
	}
}

func (s *GiftShopPart2Suite) TestFindUpperBounds() {
	tests := []struct {
		end      string
		length   int
		expected int64
	}{
		{end: "22", length: 1, expected: 2},
		{end: "111", length: 1, expected: 1},
		{end: "1012", length: 2, expected: 10},
		{end: "1110", length: 2, expected: 10},
		{end: "1188511890", length: 5, expected: 11885},
		{end: "222224", length: 3, expected: 222},
		{end: "222224", length: 1, expected: 2},
		{end: "2121212124", length: 2, expected: 21},
		{end: "115", length: 1, expected: 1},
		{end: "1012", length: 2, expected: 10},
	}
	for _, tt := range tests {
		s.Run(fmt.Sprintf("end:%s", tt.end), func() {
			actual := findUpperBoundGivenLength(tt.end, tt.length)
			s.Assert().Equal(tt.expected, actual)
		})
	}
}

func (s *GiftShopPart2Suite) TestSolvePart2Example() {
	s.Assert().Equal(int64(4174379265), SolvePart2("example_input.txt"))
}

func (s *GiftShopPart2Suite) TestSumInvalidIdsPart2() {
	tests := []struct {
		start    string
		end      string
		expected int64
	}{
		{start: "11", end: "22", expected: 33},
		{start: "95", end: "115", expected: 200},
		{start: "998", end: "1012", expected: 2009},
		{start: "1188511880", end: "1188511890", expected: 1188511885},
		{start: "222220", end: "222224", expected: 222222},
		{start: "1698522", end: "1698528", expected: 0},
	}
	for _, tt := range tests {
		s.Run(fmt.Sprintf("start[%s],end[%s],expected:[%d]", tt.start, tt.end, tt.expected), func() {
			actual := sumInvalidProductIdsForAllLengths(tt.start, tt.end)
			s.Assert().Equal(tt.expected, actual)
		})
	}
	s.Assert().Equal(int64(33), sumInvalidIds("11", "22"))
}

func TestGiftShopPart2Suite(t *testing.T) {
	suite.Run(t, new(GiftShopPart2Suite))
}
