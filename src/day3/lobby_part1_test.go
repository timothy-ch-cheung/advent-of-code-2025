package day3

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type LobbyPart1Suite struct {
	suite.Suite
}

func (l *LobbyPart1Suite) TestMaxJoltage() {
	tests := []struct {
		bank     string
		expected int
	}{
		{bank: "123", expected: 23},
		{bank: "45819", expected: 89},
	}
	for _, tt := range tests {
		l.Run("", func() {
			actual := maxJoltage(tt.bank)
			l.Assert().Equal(tt.expected, actual)
		})
	}
}

func (l *LobbyPart1Suite) TestSolvePart1Example() {
	l.Assert().Equal(357, SolvePart1("example_input.txt"))
}

func TestLobbyPart1Suite(t *testing.T) {
	suite.Run(t, new(LobbyPart1Suite))
}
