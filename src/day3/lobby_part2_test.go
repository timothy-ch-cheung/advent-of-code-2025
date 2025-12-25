package day3

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type LobbyPart2Suite struct {
	suite.Suite
}

func (l *LobbyPart2Suite) TestMaxJoltageV2() {
	tests := []struct {
		bank     string
		expected int64
	}{
		{bank: "987654321111111", expected: 987654321111},
		{bank: "811111111111119", expected: 811111111119},
		{bank: "234234234234278", expected: 434234234278},
		{bank: "818181911112111", expected: 888911112111},
	}
	for _, tt := range tests {
		l.Run("", func() {
			actual := maxJoltageV2(tt.bank, 12)
			l.Assert().Equal(tt.expected, actual)
		})
	}
}

func (l *LobbyPart2Suite) TestSolvePart2Example() {
	l.Assert().Equal(int64(3121910778619), SolvePart2("example_input.txt"))
}

func TestLobbyPart2Suite(t *testing.T) {
	suite.Run(t, new(LobbyPart2Suite))
}
