package basic_calculator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BasicCalculatorSuite struct {
	suite.Suite
}

func TestBasicCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(BasicCalculatorSuite))
}

func (s *BasicCalculatorSuite) SetupTest() {
}

func (s *BasicCalculatorSuite) Test_1() {
	assert.Equal(s.T(), 2, calculate("1 + 1"), "")
}

func (s *BasicCalculatorSuite) Test_2() {
	assert.Equal(s.T(), 6, calculate("1 + 2 + 3"), "")
}

func (s *BasicCalculatorSuite) Test_3() {
	assert.Equal(s.T(), 3, calculate(" 2-1 + 2 "), "")
}

func (s *BasicCalculatorSuite) Test_4() {
	assert.Equal(s.T(), 23, calculate("(1+(4+5+2)-3)+(6+8)"), "")
}

func (s *BasicCalculatorSuite) Test_5() {
	assert.Equal(s.T(), -12, calculate("- (3 + (4 + 5))"), "")
}

func (s *BasicCalculatorSuite) Test_6() {
	assert.Equal(s.T(), 0, calculate("1 - 1"), "")
}
