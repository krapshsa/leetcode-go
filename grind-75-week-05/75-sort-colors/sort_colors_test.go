package sort_colors

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type SortColorsSuite struct {
	suite.Suite
}

func TestSortColorsTestSuite(t *testing.T) {
	suite.Run(t, new(SortColorsSuite))
}

func (s *SortColorsSuite) SetupTest() {
}

func (s *SortColorsSuite) Test_1() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	assert.Equal(s.T(), []int{0, 0, 1, 1, 2, 2}, nums, "")
}

func (s *SortColorsSuite) Test_2() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	assert.Equal(s.T(), []int{0, 1, 2}, nums, "")
}
