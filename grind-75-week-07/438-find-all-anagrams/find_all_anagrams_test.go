package find_all_anagrams

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AnagramsSuite struct {
	suite.Suite
}

func TestAnagramsTestSuite(t *testing.T) {
	suite.Run(t, new(AnagramsSuite))
}

func (s *AnagramsSuite) SetupTest() {
}

func (s *AnagramsSuite) Test_1() {
	assert.Equal(s.T(), []int{0, 6}, findAnagrams("cbaebabacd", "abc"), "")
}

func (s *AnagramsSuite) Test_2() {
	assert.Equal(s.T(), []int{0, 1, 2}, findAnagrams("abab", "ab"), "")
}
