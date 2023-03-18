package string_to_integer

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AtoiSuite struct {
	suite.Suite
}

func TestAtoiTestSuite(t *testing.T) {
	suite.Run(t, new(AtoiSuite))
}

func (s *AtoiSuite) SetupTest() {
}

func (s *AtoiSuite) Test_1() {
	assert.Equal(s.T(), 42, myAtoi("42"), "")
}

func (s *AtoiSuite) Test_2() {
	assert.Equal(s.T(), -42, myAtoi("   -42"), "")
}

func (s *AtoiSuite) Test_3() {
	assert.Equal(s.T(), 4193, myAtoi("4193 with words"), "")
}

func (s *AtoiSuite) Test_4() {
	assert.Equal(s.T(), 2147483647, myAtoi("21474836477"), "")
}

func (s *AtoiSuite) Test_5() {
	assert.Equal(s.T(), -2147483648, myAtoi("-21474836481"), "")
}

func (s *AtoiSuite) Test_6() {
	assert.Equal(s.T(), 0, myAtoi("words and 987"), "")
}

func (s *AtoiSuite) Test_7() {
	assert.Equal(s.T(), 3, myAtoi("3.14159"), "")
}

func (s *AtoiSuite) Test_8() {
	assert.Equal(s.T(), 12345678, myAtoi("  0000000000012345678"), "")
}

func (s *AtoiSuite) Test_9() {
	assert.Equal(s.T(), 0, myAtoi("00000-42a1234"), "")
}
