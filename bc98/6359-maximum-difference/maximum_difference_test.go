package maximum_difference

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MaximumDifferenceSuite struct {
	suite.Suite
}

func TestMaximumDifferenceTestSuite(t *testing.T) {
	suite.Run(t, new(MaximumDifferenceSuite))
}

func (s *MaximumDifferenceSuite) SetupTest() {
}

func (s *MaximumDifferenceSuite) Test_1() {
	assert.Equal(s.T(), 99009, minMaxDifference(11891), "")
}

func (s *MaximumDifferenceSuite) Test_2() {
	assert.Equal(s.T(), 99, minMaxDifference(90), "")
}

func (s *MaximumDifferenceSuite) Test_3() {
	assert.Equal(s.T(), 900, minMaxDifference(456), "")
}
