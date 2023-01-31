package search_in_rotated_sorted_array

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type RotatedSuite struct {
	suite.Suite
}

func TestRotatedTestSuite(t *testing.T) {
	suite.Run(t, new(RotatedSuite))
}

func (s *RotatedSuite) SetupTest() {
}

func (s *RotatedSuite) Test_1() {
	assert.Equal(s.T(), 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0), "")
}

func (s *RotatedSuite) Test_2() {
	assert.Equal(s.T(), -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3), "")
}

func (s *RotatedSuite) Test_3() {
	assert.Equal(s.T(), -1, search([]int{1}, 0), "")
}
