package largest_rectangle

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LargetRectangleSuite struct {
	suite.Suite
}

func TestLargetRectangleTestSuite(t *testing.T) {
	suite.Run(t, new(LargetRectangleSuite))
}

func (s *LargetRectangleSuite) SetupTest() {
}

func (s *LargetRectangleSuite) Test_1() {
	assert.Equal(s.T(), 10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}), "")
}

func (s *LargetRectangleSuite) Test_2() {
	assert.Equal(s.T(), 4, largestRectangleArea([]int{2, 4}), "")
}

func (s *LargetRectangleSuite) Test_3() {
	assert.Equal(s.T(), 2, largestRectangleArea([]int{1, 1}), "")
}

func (s *LargetRectangleSuite) Test_4() {
	assert.Equal(s.T(), 20, largestRectangleArea([]int{3, 6, 5, 7, 4, 8, 1, 0}), "")
}

func (s *LargetRectangleSuite) Test_5() {
	assert.Equal(s.T(), 24, largestRectangleArea([]int{3, 5, 5, 2, 5, 5, 6, 6, 4, 4, 1, 1, 2, 5, 5, 6, 6, 4, 1, 3}), "")
}
