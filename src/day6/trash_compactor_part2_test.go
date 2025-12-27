package day6

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TrashCompactorPart2Suite struct {
	suite.Suite
}

func (t *TrashCompactorPart2Suite) TestSolvePart2() {
	t.Assert().Equal(3263827, SolvePart2("example_input.txt"))
}

func TestTrashCompactorPart2SuiteTest(t *testing.T) {
	suite.Run(t, new(TrashCompactorPart2Suite))
}
