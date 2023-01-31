package number_of_islands

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type IslandSuite struct {
	suite.Suite
}

func TestIslandTestSuite(t *testing.T) {
	suite.Run(t, new(IslandSuite))
}

func (s *IslandSuite) SetupTest() {
}

func (s *IslandSuite) Test_1() {
	grid := [][]byte{
		{byte('1'), byte('1'), byte('1'), byte('1'), byte('0')},
		{byte('1'), byte('1'), byte('0'), byte('1'), byte('0')},
		{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
		{byte('0'), byte('0'), byte('0'), byte('0'), byte('0')},
	}
	assert.Equal(s.T(), 1, numIslands(grid), "")
}

func (s *IslandSuite) Test_2() {
	grid := [][]byte{
		{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
		{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
		{byte('0'), byte('0'), byte('1'), byte('0'), byte('0')},
		{byte('0'), byte('0'), byte('0'), byte('1'), byte('1')},
	}
	assert.Equal(s.T(), 3, numIslands(grid), "")
}
