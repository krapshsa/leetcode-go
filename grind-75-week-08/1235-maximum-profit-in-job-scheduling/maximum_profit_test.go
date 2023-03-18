package maximum_profit_in_job_scheduling

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MaximumProfitSuite struct {
	suite.Suite
}

func TestMaximumProfitTestSuite(t *testing.T) {
	suite.Run(t, new(MaximumProfitSuite))
}

func (s *MaximumProfitSuite) SetupTest() {
}

func (s *MaximumProfitSuite) Test_1() {
	assert.Equal(s.T(), 120, jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}), "")
}

func (s *MaximumProfitSuite) Test_2() {
	assert.Equal(s.T(), 150, jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}), "")
}

func (s *MaximumProfitSuite) Test_3() {
	assert.Equal(s.T(), 6, jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}), "")
}

func (s *MaximumProfitSuite) Test_4() {
	assert.Equal(s.T(), 7, jobScheduling([]int{1, 2, 2, 3}, []int{2, 5, 3, 4}, []int{3, 4, 1, 2}), "")
}
