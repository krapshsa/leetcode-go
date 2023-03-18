package minimum_window_substring

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MWSSuite struct {
	suite.Suite
}

func TestMWSTestSuite(t *testing.T) {
	suite.Run(t, new(MWSSuite))
}

func (s *MWSSuite) SetupTest() {
}

func (s *MWSSuite) Test_1() {
	assert.Equal(s.T(), "BANC", minWindow("ADOBECODEBANC", "ABC"), "")
}

func (s *MWSSuite) Test_2() {
	assert.Equal(s.T(), "a", minWindow("a", "a"), "")
}

func (s *MWSSuite) Test_3() {
	assert.Equal(s.T(), "", minWindow("a", "aa"), "")
}

func (s *MWSSuite) Test_4() {
	assert.Equal(s.T(), "abbbbbcdd", minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"), "")
}
