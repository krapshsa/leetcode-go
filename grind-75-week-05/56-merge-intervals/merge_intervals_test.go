package merge_intervals

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type MergeIntervalsSuite struct {
	suite.Suite
}

func TestMergeIntervalsTestSuite(t *testing.T) {
	suite.Run(t, new(MergeIntervalsSuite))
}

func (s *MergeIntervalsSuite) SetupTest() {
}

func (s *MergeIntervalsSuite) Test_1() {
	assert.Equal(s.T(), [][]int{{1, 6}, {8, 10}, {15, 18}}, merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}), "")
}

func (s *MergeIntervalsSuite) Test_2() {
	assert.Equal(s.T(), [][]int{{1, 5}}, merge([][]int{{1, 4}, {4, 5}}), "")
}
