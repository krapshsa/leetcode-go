package square_free_subsets

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SqureFreeSuite struct {
	suite.Suite
}

func TestSqureFreeTestSuite(t *testing.T) {
	suite.Run(t, new(SqureFreeSuite))
}

func (s *SqureFreeSuite) SetupTest() {
}

func (s *SqureFreeSuite) Test_1() {
	assert.Equal(s.T(), 3, squareFreeSubsets([]int{3, 4, 4, 5}), "")
}

func (s *SqureFreeSuite) Test_2() {
	assert.Equal(s.T(), 1, squareFreeSubsets([]int{1}), "")
}

func (s *SqureFreeSuite) Test_3() {
	assert.Equal(s.T(), 1, squareFreeSubsets([]int{4, 6, 6, 18, 26, 27}), "")
}
