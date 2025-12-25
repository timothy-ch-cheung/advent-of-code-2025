package day4

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PrintingDepartmentPart2Suite struct {
	suite.Suite
}

func (p *PrintingDepartmentPart2Suite) TestSolvePart2() {
	p.Assert().Equal(43, SolvePart2("example_input.txt"))
}

func TestPrintDepartmentPart2(t *testing.T) {
	suite.Run(t, new(PrintingDepartmentPart2Suite))
}
