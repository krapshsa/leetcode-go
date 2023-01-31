package lca_of_a_binary_tree

import (
	"testing"
)
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type LCASuite struct {
	suite.Suite
}

func TestLCATestSuite(t *testing.T) {
	suite.Run(t, new(LCASuite))
}

func (s *LCASuite) SetupTest() {
}

func (s *LCASuite) Test_1() {
	p := &TreeNode{
		5,
		&TreeNode{6, nil, nil},
		&TreeNode{
			2,
			&TreeNode{7, nil, nil},
			&TreeNode{4, nil, nil}},
	}
	q := &TreeNode{
		1,
		&TreeNode{0, nil, nil},
		&TreeNode{8, nil, nil},
	}
	root := &TreeNode{3, p, q}
	assert.Equal(s.T(), root, lowestCommonAncestor(root, p, q), "")
}

func (s *LCASuite) Test_2() {
	q := &TreeNode{4, nil, nil}
	p := &TreeNode{
		5,
		&TreeNode{6, nil, nil},
		&TreeNode{
			2,
			&TreeNode{7, nil, nil},
			q,
		},
	}
	root := &TreeNode{
		3,
		p,
		&TreeNode{
			1,
			&TreeNode{0, nil, nil},
			&TreeNode{8, nil, nil},
		},
	}
	assert.Equal(s.T(), p, lowestCommonAncestor(root, p, q), "")
}

func (s *LCASuite) Test_3() {
	q := &TreeNode{2, nil, nil}
	root := &TreeNode{1, q, nil}
	assert.Equal(s.T(), root, lowestCommonAncestor(root, root, q), "")
}
