package minimum_operations

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MinOperationsSuite struct {
	suite.Suite
}

func TestMinOperationsTestSuite(t *testing.T) {
	suite.Run(t, new(MinOperationsSuite))
}

func (s *MinOperationsSuite) SetupTest() {
}

func (s *MinOperationsSuite) Test_1() {
	assert.Equal(s.T(), 3, minOperations(39), "")
}

func (s *MinOperationsSuite) Test_2() {
	assert.Equal(s.T(), 3, minOperations(54), "")
}

func (s *MinOperationsSuite) Test_3() {
	assert.Equal(s.T(), 2, minOperations(9), "")
}

func (s *MinOperationsSuite) Test_4() {
	assert.Equal(s.T(), 3, minOperations(27), "")
}
