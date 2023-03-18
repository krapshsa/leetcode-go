package LRU_cache

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LRUCacheSuite struct {
	suite.Suite
}

func TestLRUCacheTestSuite(t *testing.T) {
	suite.Run(t, new(LRUCacheSuite))
}

func (s *LRUCacheSuite) SetupTest() {
}

func (s *LRUCacheSuite) Test_1() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	assert.Equal(s.T(), 1, obj.Get(1), "return 1")
	obj.Put(3, 3)
	assert.Equal(s.T(), -1, obj.Get(2), "returns -1 (not found)")
	obj.Put(4, 4)
	assert.Equal(s.T(), -1, obj.Get(1), "return -1 (not found)")
	assert.Equal(s.T(), 3, obj.Get(3), "return 3")
	assert.Equal(s.T(), 4, obj.Get(4), "return 4")
}

func (s *LRUCacheSuite) Test_2() {
	obj := Constructor(2)
	obj.Put(1, 0)
	obj.Put(2, 2)
	assert.Equal(s.T(), 0, obj.Get(1), "return 0")
	obj.Put(3, 3)
	assert.Equal(s.T(), -1, obj.Get(2), "returns -1 (not found)")
	obj.Put(4, 4)
	assert.Equal(s.T(), -1, obj.Get(1), "return -1 (not found)")
	assert.Equal(s.T(), 3, obj.Get(3), "return 3")
	assert.Equal(s.T(), 4, obj.Get(4), "return 4")
}

func (s *LRUCacheSuite) Test_3() {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	assert.Equal(s.T(), -1, obj.Get(1), "returns -1 (not found)")
	assert.Equal(s.T(), 3, obj.Get(2), "reutnr 3")
}
