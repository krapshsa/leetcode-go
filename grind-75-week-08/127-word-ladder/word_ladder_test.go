package word_ladder

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WordLadderSuite struct {
	suite.Suite
}

func TestWordLadderTestSuite(t *testing.T) {
	suite.Run(t, new(WordLadderSuite))
}

func (s *WordLadderSuite) SetupTest() {
}

func (s *WordLadderSuite) Test_1() {
	strings := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	assert.Equal(s.T(), 5, ladderLength("hit", "cog", strings), "")
}

func (s *WordLadderSuite) Test_2() {
	strings := []string{"hot", "dot", "dog", "lot", "log"}
	assert.Equal(s.T(), 0, ladderLength("hit", "cog", strings), "")
}
