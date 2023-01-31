package implement_trie

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type TrieSuite struct {
	suite.Suite
}

func TestTrieTestSuite(t *testing.T) {
	suite.Run(t, new(TrieSuite))
}

func (s *TrieSuite) SetupTest() {
}

func (s *TrieSuite) Test_1() {
	obj := Constructor()
	obj.Insert("apple")
	assert.Equal(s.T(), true, obj.Search("apple"), "")
	assert.Equal(s.T(), false, obj.Search("app"), "")
	assert.Equal(s.T(), true, obj.StartsWith("app"), "")

	obj.Insert("app")
	assert.Equal(s.T(), true, obj.Search("app"), "")
}
