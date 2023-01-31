package account_merge

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type AccountMergeSuite struct {
	suite.Suite
}

func TestAccountMergeTestSuite(t *testing.T) {
	suite.Run(t, new(AccountMergeSuite))
}

func (s *AccountMergeSuite) SetupTest() {
}

func (s *AccountMergeSuite) Test_1() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}
	expected := [][]string{
		{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}
	assert.Equal(s.T(), expected, accountsMerge(accounts), "")
}

func (s *AccountMergeSuite) Test_2() {
	accounts := [][]string{
		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
		{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
		{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
		{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
		{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
	}
	expected := [][]string{
		{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
		{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
		{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
		{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
		{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
	}
	assert.Equal(s.T(), expected, accountsMerge(accounts), "")
}

func (s *AccountMergeSuite) Test_3() {
	accounts := [][]string{
		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
		{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
		{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
		{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
		{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
	}
	expected := [][]string{
		{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
		{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
		{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
		{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
		{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
	}
	assert.Equal(s.T(), expected, accountsMerge(accounts), "")
}
