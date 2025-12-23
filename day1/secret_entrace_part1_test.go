package day1

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SecretEntranceSuite struct {
	suite.Suite
}

func (s *SecretEntranceSuite) TestRotateLeft() {
	s.Assert().Equal(2, rotateLeft(50, 48))
}

func (s *SecretEntranceSuite) TestRotateUnderflow() {
	s.Assert().Equal(99, rotateLeft(2, 3))
}

func (s *SecretEntranceSuite) TestRotateDoubleUnderflow() {
	s.Assert().Equal(99, rotateLeft(2, 103))
}

func (s *SecretEntranceSuite) TestRotateRight() {
	s.Assert().Equal(52, rotateRight(50, 2))
}

func (s *SecretEntranceSuite) TestRotateRightOverflow() {
	s.Assert().Equal(2, rotateRight(98, 4))
}

func (s *SecretEntranceSuite) TestRotateRightDoubleOverflow() {
	s.Assert().Equal(2, rotateRight(98, 104))
}

func TestSecretEntranceSuite(t *testing.T) {
	suite.Run(t, new(SecretEntranceSuite))
}
