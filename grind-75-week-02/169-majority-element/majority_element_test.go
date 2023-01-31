package majority_element

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type Majority struct {
	suite.Suite
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(Majority))
}

func (s *Majority) SetupTest() {
}

func (s *Majority) Test_1() {
	assert.Equal(s.T(), 3, majorityElement([]int{3, 2, 3}), "")
}

func (s *Majority) Test_2() {
	assert.Equal(s.T(), 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}), "")
}
