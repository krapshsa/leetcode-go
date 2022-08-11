package diameter_of_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{3, nil, nil},
	}
	assert.Equal(t, 3, diameterOfBinaryTree(tree))
}

func Test_diameterOfBinaryTree_2(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		nil,
	}
	assert.Equal(t, 1, diameterOfBinaryTree(tree))
}
