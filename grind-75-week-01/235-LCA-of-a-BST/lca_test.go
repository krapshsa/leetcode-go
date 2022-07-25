package LCA_of_a_BST

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	p := &TreeNode{
		2,
		&TreeNode{0, nil, nil},
		&TreeNode{
			4,
			&TreeNode{3, nil, nil},
			&TreeNode{5, nil, nil},
		},
	}
	q := &TreeNode{
		8,
		&TreeNode{7, nil, nil},
		&TreeNode{9, nil, nil},
	}
	root := &TreeNode{6, p, q}
	assert.Equal(t, 6, lowestCommonAncestor(root, p, q).Val)
}

func Test_lowestCommonAncestor_1(t *testing.T) {
	q := &TreeNode{
		4,
		&TreeNode{3, nil, nil},
		&TreeNode{5, nil, nil},
	}
	p := &TreeNode{
		2,
		&TreeNode{0, nil, nil},
		q,
	}
	root := &TreeNode{
		6,
		p,
		&TreeNode{
			8,
			&TreeNode{7, nil, nil},
			&TreeNode{9, nil, nil},
		},
	}
	assert.Equal(t, 2, lowestCommonAncestor(root, p, q).Val)
}

func Test_lowestCommonAncestor_2(t *testing.T) {
	p := &TreeNode{1, nil, nil}
	q := &TreeNode{2, p, nil}
	root := q
	assert.Equal(t, 2, lowestCommonAncestor(root, p, q).Val)
}
