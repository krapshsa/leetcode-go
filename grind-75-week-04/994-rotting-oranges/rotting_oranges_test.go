package rotting_oranges

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type RottingOrangesSuite struct {
	suite.Suite
}

func TestRottingOrangesTestSuite(t *testing.T) {
	suite.Run(t, new(RottingOrangesSuite))
}

func (s *RottingOrangesSuite) SetupTest() {
}

func (s *RottingOrangesSuite) Test_1() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	assert.Equal(s.T(), 4, orangesRotting(grid), "")
}

func (s *RottingOrangesSuite) Test_2() {
	grid := [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}
	assert.Equal(s.T(), -1, orangesRotting(grid), "")
}

func (s *RottingOrangesSuite) Test_3() {
	grid := [][]int{{0, 2}}
	assert.Equal(s.T(), 0, orangesRotting(grid), "")
}
