package left_and_right_sum_diff

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LeftRightSuite struct {
	suite.Suite
}

func TestLeftRightTestSuite(t *testing.T) {
	suite.Run(t, new(LeftRightSuite))
}

func (s *LeftRightSuite) SetupTest() {
}

func (s *LeftRightSuite) Test_1() {
	assert.Equal(s.T(), []int{15, 1, 11, 22}, leftRigthDifference([]int{10, 4, 8, 3}), "")
}
