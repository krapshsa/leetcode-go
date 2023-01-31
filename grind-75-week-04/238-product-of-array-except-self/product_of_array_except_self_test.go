package product_of_array_except_self

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type ExceptSelfSuite struct {
	suite.Suite
}

func TestExceptSelfTestSuite(t *testing.T) {
	suite.Run(t, new(ExceptSelfSuite))
}

func (s *ExceptSelfSuite) SetupTest() {
}

func (s *ExceptSelfSuite) Test_1() {
	assert.Equal(s.T(), []int{24, 12, 8, 6}, productExceptSelf([]int{1, 2, 3, 4}), "")
}
