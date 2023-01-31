package ongest_substring

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type LongestSuite struct {
	suite.Suite
}

func TestLongestTestSuite(t *testing.T) {
	suite.Run(t, new(LongestSuite))
}

func (s *LongestSuite) SetupTest() {
}

func (s *LongestSuite) Test_1() {
	assert.Equal(s.T(), 3, lengthOfLongestSubstring("abcabcbb"), "The answer is \"abc\", with the length of 3.")
}

func (s *LongestSuite) Test_2() {
	assert.Equal(s.T(), 1, lengthOfLongestSubstring("bbbbb"), "The answer is \"b\", with the length of 1.")
}

func (s *LongestSuite) Test_3() {
	assert.Equal(s.T(), 3, lengthOfLongestSubstring("pwwkew"), "The answer is \"wke\", with the length of 3.\nNotice that the answer must be a substring, \"pwke\" is a subsequence and not a substring.")
}

func (s *LongestSuite) Test_4() {
	assert.Equal(s.T(), 3, lengthOfLongestSubstring("dvdf"), "")
}
