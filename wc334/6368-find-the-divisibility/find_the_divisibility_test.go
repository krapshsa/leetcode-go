package find_the_divisibility

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DivisibilitySuite struct {
	suite.Suite
}

func TestDivisibilityTestSuite(t *testing.T) {
	suite.Run(t, new(DivisibilitySuite))
}

func (s *DivisibilitySuite) SetupTest() {
}

func (s *DivisibilitySuite) Test_1() {
	assert.Equal(s.T(), []int{1, 1, 0, 0, 0, 1, 1, 0, 0}, divisibilityArray("998244353", 3), "")
}

func (s *DivisibilitySuite) Test_2() {
	assert.Equal(s.T(), []int{0, 1, 0, 1}, divisibilityArray("1010", 10), "")
}

func (s *DivisibilitySuite) Test_3() {
	assert.Equal(s.T(), []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, divisibilityArray("91221181269244172125025075166510211202115152121212341281327", 21), "")
}
