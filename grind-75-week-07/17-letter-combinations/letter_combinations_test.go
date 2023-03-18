package letter_combinations

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LetterCombinationsSuite struct {
	suite.Suite
}

func TestLetterCombinationsTestSuite(t *testing.T) {
	suite.Run(t, new(LetterCombinationsSuite))
}

func (s *LetterCombinationsSuite) SetupTest() {
}

func (s *LetterCombinationsSuite) Test_1() {
	assert.Equal(s.T(), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"), "")
}

func (s *LetterCombinationsSuite) Test_2() {
	assert.Equal(s.T(), []string{"a", "b", "c"}, letterCombinations("2"), "")
}
