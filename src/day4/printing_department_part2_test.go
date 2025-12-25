package day4

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PrintDepartmentPart2Suite struct {
	suite.Suite
}

func TestPrintDepartmentPart2(t *testing.T) {
	suite.Run(t, new(PrintDepartmentPart2Suite))
}
