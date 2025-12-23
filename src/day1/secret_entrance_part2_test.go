package day1

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SecretEntrancePart2Suite struct {
	suite.Suite
}

func (s *SecretEntrancePart2Suite) TestRotateLeft() {
	result := rotateLeftV2(50, 48)
	s.Assert().Equal(2, result.newPos)
	s.Assert().Equal(0, result.clicks)
}

func (s *SecretEntrancePart2Suite) TestRotateLeftUnderflow() {
	result := rotateLeftV2(2, 3)
	s.Assert().Equal(99, result.newPos)
	s.Assert().Equal(1, result.clicks)
}

func (s *SecretEntrancePart2Suite) TestRotateLeftDoubleUnderflow() {
	result := rotateLeftV2(2, 103)
	s.Assert().Equal(99, result.newPos)
	s.Assert().Equal(2, result.clicks)
}

func (s *SecretEntrancePart2Suite) TestRotateRight() {
	result := rotateRightV2(50, 2)
	s.Assert().Equal(52, result.newPos)
	s.Assert().Equal(0, result.clicks)
}

func (s *SecretEntrancePart2Suite) TestRotateRightOverflow() {
	result := rotateRightV2(98, 4)
	s.Assert().Equal(2, result.newPos)
	s.Assert().Equal(1, result.clicks)
}

func (s *SecretEntrancePart2Suite) TestRotateRightDoubleOverflow() {
	result := rotateRightV2(98, 104)
	s.Assert().Equal(2, result.newPos)
	s.Assert().Equal(2, result.clicks)
}

func (s *SecretEntrancePart2Suite) TestRotate1000() {
	result := rotateRightV2(50, 1000)
	s.Assert().Equal(50, result.newPos)
	s.Assert().Equal(10, result.clicks)
}

func (s *SecretEntrancePart1Suite) TestSolvePart2ExampleStepByStep() {
	result1 := rotateLeftV2(50, 68)
	s.Assert().Equal(82, result1.newPos)
	s.Assert().Equal(1, result1.clicks)

	result2 := rotateLeftV2(82, 30)
	s.Assert().Equal(52, result2.newPos)
	s.Assert().Equal(0, result2.clicks)

	result3 := rotateRightV2(52, 48)
	s.Assert().Equal(0, result3.newPos)
	s.Assert().Equal(0, result3.clicks)

	result4 := rotateLeftV2(0, 5)
	s.Assert().Equal(95, result4.newPos)
	s.Assert().Equal(0, result4.clicks)

	result5 := rotateRightV2(95, 60)
	s.Assert().Equal(55, result5.newPos)
	s.Assert().Equal(1, result5.clicks)

	result6 := rotateLeftV2(55, 55)
	s.Assert().Equal(0, result6.newPos)
	s.Assert().Equal(0, result6.clicks)

	result7 := rotateLeftV2(0, 1)
	s.Assert().Equal(99, result7.newPos)
	s.Assert().Equal(0, result7.clicks)

	result8 := rotateLeftV2(99, 99)
	s.Assert().Equal(0, result8.newPos)
	s.Assert().Equal(0, result8.clicks)

	result9 := rotateRightV2(0, 14)
	s.Assert().Equal(14, result9.newPos)
	s.Assert().Equal(0, result9.clicks)

	result10 := rotateLeftV2(14, 82)
	s.Assert().Equal(32, result10.newPos)
	s.Assert().Equal(1, result10.clicks)
}

func (s *SecretEntrancePart1Suite) TestSolvePart2Example() {
	s.Assert().Equal(6, SolvePart2("example_input.txt"))
}

func TestSecretEntrancePart2Suite(t *testing.T) {
	suite.Run(t, new(SecretEntrancePart2Suite))
}
