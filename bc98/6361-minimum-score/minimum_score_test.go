package minimum_score

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MimimumScoreSuite struct {
	suite.Suite
}

func TestMimimumScoreTestSuite(t *testing.T) {
	suite.Run(t, new(MimimumScoreSuite))
}

func (s *MimimumScoreSuite) SetupTest() {
}

func (s *MimimumScoreSuite) Test_1() {
	assert.Equal(s.T(), 0, minimizeSum([]int{1, 4, 3}), "")
}

func (s *MimimumScoreSuite) Test_2() {
	assert.Equal(s.T(), 3, minimizeSum([]int{1, 4, 7, 8, 5}), "")
}

func (s *MimimumScoreSuite) Test_3() {
	assert.Equal(s.T(), 24, minimizeSum([]int{59, 27, 9, 81, 33}), "")
}

func (s *MimimumScoreSuite) Test_4() {
	assert.Equal(s.T(), 30, minimizeSum([]int{8, 28, 42, 58, 75}), "")
}

func (s *MimimumScoreSuite) Test_5() {
	assert.Equal(s.T(), 14, minimizeSum([]int{31, 25, 72, 79, 74, 65}), "")
}
