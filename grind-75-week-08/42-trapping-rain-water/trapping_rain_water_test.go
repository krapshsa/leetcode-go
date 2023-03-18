package trapping_rain_water

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TrappingSuite struct {
	suite.Suite
}

func TestTrappingTestSuite(t *testing.T) {
	suite.Run(t, new(TrappingSuite))
}

func (s *TrappingSuite) SetupTest() {
}

func (s *TrappingSuite) Test_1() {
	assert.Equal(s.T(), 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), "")
}

func (s *TrappingSuite) Test_2() {
	assert.Equal(s.T(), 9, trap([]int{4, 2, 0, 3, 2, 5}), "")
}
