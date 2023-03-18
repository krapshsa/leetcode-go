package maximum_number_marked

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MaxNumMarkedSuite struct {
	suite.Suite
}

func TestMaxNumMarkedTestSuite(t *testing.T) {
	suite.Run(t, new(MaxNumMarkedSuite))
}

func (s *MaxNumMarkedSuite) SetupTest() {
}

func (s *MaxNumMarkedSuite) Test_1() {
	assert.Equal(s.T(), 2, maxNumOfMarkedIndices([]int{3, 5, 2, 4}), "")
}

func (s *MaxNumMarkedSuite) Test_2() {
	assert.Equal(s.T(), 4, maxNumOfMarkedIndices([]int{9, 2, 5, 4}), "")
}

func (s *MaxNumMarkedSuite) Test_3() {
	assert.Equal(s.T(), 0, maxNumOfMarkedIndices([]int{7, 6, 8}), "")
}

func (s *MaxNumMarkedSuite) Test_4() {
	assert.Equal(s.T(), 26, maxNumOfMarkedIndices([]int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40}), "")
}
