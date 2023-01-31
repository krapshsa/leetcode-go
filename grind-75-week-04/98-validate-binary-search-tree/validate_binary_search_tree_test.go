package validate_binary_search_tree

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type ValidateBSTSuite struct {
	suite.Suite
}

func TestValidateBSTTestSuite(t *testing.T) {
	suite.Run(t, new(ValidateBSTSuite))
}

func (s *ValidateBSTSuite) SetupTest() {
}

func (s *ValidateBSTSuite) Test_1() {
	root := &TreeNode{
		2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, nil, nil},
	}
	assert.Equal(s.T(), true, isValidBST(root), "")
}

func (s *ValidateBSTSuite) Test_2() {
	root := &TreeNode{
		5,
		&TreeNode{1, nil, nil},
		&TreeNode{
			4,
			&TreeNode{3, nil, nil},
			&TreeNode{6, nil, nil},
		},
	}
	assert.Equal(s.T(), false, isValidBST(root), "")
}
