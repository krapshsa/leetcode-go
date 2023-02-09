package word_break

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type WordBreakSuite struct {
	suite.Suite
}

func TestWordBreakTestSuite(t *testing.T) {
	suite.Run(t, new(WordBreakSuite))
}

func (s *WordBreakSuite) SetupTest() {
}

func (s *WordBreakSuite) Test_1() {
	assert.Equal(s.T(), true, wordBreak("leetcode", []string{"leet", "code"}), "")
}

func (s *WordBreakSuite) Test_2() {
	assert.Equal(s.T(), true, wordBreak("applepenapple", []string{"apple", "pen"}), "")
}

func (s *WordBreakSuite) Test_3() {
	assert.Equal(s.T(), false, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), "")
}
