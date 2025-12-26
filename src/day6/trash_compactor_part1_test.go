package day6

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TrashCompactorPart1Suite struct {
	suite.Suite
}

func (t *TrashCompactorPart1Suite) TestTrashCompactorPart1() {
	t.Assert().Equal(4277556, SolvePart1("example_input.txt"))
}

func TestTrashCompactorPart1Suite(t *testing.T) {
	suite.Run(t, new(TrashCompactorPart1Suite))
}
