package find_median_from_data_stream

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FindMedianSuite struct {
	suite.Suite
}

func TestFindMedianTestSuite(t *testing.T) {
	suite.Run(t, new(FindMedianSuite))
}

func (s *FindMedianSuite) SetupTest() {
}

func (s *FindMedianSuite) Test_1() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	assert.Equal(s.T(), 1.5, obj.FindMedian(), "")
	obj.AddNum(3)
	assert.Equal(s.T(), 2.0, obj.FindMedian(), "")
}

func (s *FindMedianSuite) Test_2() {
	obj := Constructor()
	obj.AddNum(-1)
	assert.Equal(s.T(), -1.0, obj.FindMedian(), "")
	obj.AddNum(-2)
	assert.Equal(s.T(), -1.5, obj.FindMedian(), "")
	obj.AddNum(-3)
	assert.Equal(s.T(), -2.0, obj.FindMedian(), "")
	obj.AddNum(-4)
	assert.Equal(s.T(), -2.5, obj.FindMedian(), "")
	obj.AddNum(-5)
	assert.Equal(s.T(), -3.0, obj.FindMedian(), "")
}
