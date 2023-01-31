package time_based_key_value_store

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type KVSuite struct {
	suite.Suite
}

func TestKVTestSuite(t *testing.T) {
	suite.Run(t, new(KVSuite))
}

func (s *KVSuite) SetupTest() {
}

func (s *KVSuite) Test_1() {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)

	result1 := timeMap.Get("foo", 1)
	assert.Equal(s.T(), "bar", result1, "")

	result3 := timeMap.Get("foo", 3)
	assert.Equal(s.T(), "bar", result3, "")

	timeMap.Set("foo", "bar2", 4)

	result4 := timeMap.Get("foo", 4)
	assert.Equal(s.T(), "bar2", result4, "")

	result5 := timeMap.Get("foo", 5)
	assert.Equal(s.T(), "bar2", result5, "")
}
