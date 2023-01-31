package min_stack

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type MinStackSuite struct {
	suite.Suite
}

func TestMinStackTestSuite(t *testing.T) {
	suite.Run(t, new(MinStackSuite))
}

func (s *MinStackSuite) SetupTest() {
}

func (s *MinStackSuite) Test_1() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	assert.Equal(s.T(), -3, obj.GetMin(), "")
	obj.Pop()
	assert.Equal(s.T(), 0, obj.Top(), "")
	assert.Equal(s.T(), -2, obj.GetMin(), "")
}
