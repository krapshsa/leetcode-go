package permutations

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type PermutationsSuite struct {
	suite.Suite
}

func TestPermutationsTestSuite(t *testing.T) {
	suite.Run(t, new(PermutationsSuite))
}

func (s *PermutationsSuite) SetupTest() {
}

func (s *PermutationsSuite) Test_1() {
	assert.Equal(s.T(), [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute([]int{1, 2, 3}), "")
}

func (s *PermutationsSuite) Test_2() {
	assert.Equal(s.T(), [][]int{{0, 1}, {1, 0}}, permute([]int{0, 1}), "")
}

func (s *PermutationsSuite) Test_3() {
	assert.Equal(s.T(), [][]int{{1}}, permute([]int{1}), "")
}
