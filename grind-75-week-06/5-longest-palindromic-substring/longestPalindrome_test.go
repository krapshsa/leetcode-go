package longest_palindromic_substring

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PalindromeSuite struct {
	suite.Suite
}

func TestPalindromeTestSuite(t *testing.T) {
	suite.Run(t, new(PalindromeSuite))
}

func (s *PalindromeSuite) SetupTest() {
}

func (s *PalindromeSuite) Test_1() {
	assert.Equal(s.T(), "bab", longestPalindrome("babad"), "")
}

func (s *PalindromeSuite) Test_2() {
	assert.Equal(s.T(), "bb", longestPalindrome("cbbd"), "")
}
