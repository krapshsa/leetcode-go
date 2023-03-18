package merge_two_2d_array

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MergeArraySuite struct {
	suite.Suite
}

func TestMergeArrayTestSuite(t *testing.T) {
	suite.Run(t, new(MergeArraySuite))
}

func (s *MergeArraySuite) SetupTest() {
}

func (s *MergeArraySuite) Test_1() {
	assert.Equal(
		s.T(),
		[][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		mergeArrays([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}}),
		"",
	)
}

func (s *MergeArraySuite) Test_2() {
	assert.Equal(
		s.T(),
		[][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}},
		mergeArrays([][]int{{2, 4}, {3, 6}, {5, 5}}, [][]int{{1, 3}, {4, 3}}),
		"",
	)
}
