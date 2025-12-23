package day1

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SecretEntrancePart1Suite struct {
	suite.Suite
}

func (s *SecretEntrancePart1Suite) TestRotateLeft() {
	s.Assert().Equal(2, rotateLeft(50, 48))
}

func (s *SecretEntrancePart1Suite) TestRotateLeftUnderflow() {
	s.Assert().Equal(99, rotateLeft(2, 3))
}

func (s *SecretEntrancePart1Suite) TestRotateLeftDoubleUnderflow() {
	s.Assert().Equal(99, rotateLeft(2, 103))
}

func (s *SecretEntrancePart1Suite) TestRotateRight() {
	s.Assert().Equal(52, rotateRight(50, 2))
}

func (s *SecretEntrancePart1Suite) TestRotateRightOverflow() {
	s.Assert().Equal(2, rotateRight(98, 4))
}

func (s *SecretEntrancePart1Suite) TestRotateRightDoubleOverflow() {
	s.Assert().Equal(2, rotateRight(98, 104))
}

func (s *SecretEntrancePart1Suite) TestSolvePart1Example() {
	s.Assert().Equal(3, SolvePart1("example_input.txt"))
}

func TestSecretEntrancePart1Suite(t *testing.T) {
	suite.Run(t, new(SecretEntrancePart1Suite))
}
