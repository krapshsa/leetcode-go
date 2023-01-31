package coin_change

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type CoinChangeSuite struct {
	suite.Suite
}

func TestCoinChangeTestSuite(t *testing.T) {
	suite.Run(t, new(CoinChangeSuite))
}

func (s *CoinChangeSuite) SetupTest() {
}

func (s *CoinChangeSuite) Test_1() {
	assert.Equal(s.T(), 3, coinChange([]int{1, 2, 5}, 11), "")
}

func (s *CoinChangeSuite) Test_2() {
	assert.Equal(s.T(), -1, coinChange([]int{2}, 3), "")
}

func (s *CoinChangeSuite) Test_3() {
	assert.Equal(s.T(), 0, coinChange([]int{1}, 0), "")
}
