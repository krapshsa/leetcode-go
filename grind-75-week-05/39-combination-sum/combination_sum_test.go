package combination_sum

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CombinationSumSuite struct {
	suite.Suite
}

func TestCombinationSumTestSuite(t *testing.T) {
	suite.Run(t, new(CombinationSumSuite))
}

func (s *CombinationSumSuite) SetupTest() {
}

func (s *CombinationSumSuite) Test_1() {
	assert.Equal(s.T(), [][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7), "")
}

func (s *CombinationSumSuite) Test_2() {
	assert.Equal(s.T(), [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8), "")
}

func (s *CombinationSumSuite) Test_3() {
	assert.Equal(s.T(), [][]int{}, combinationSum([]int{2}, 1), "")
}
