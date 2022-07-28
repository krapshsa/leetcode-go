package balanced_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isBalanced_true(t *testing.T) {
	root := &TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}},
	}
	assert.Equal(t, true, isBalanced(root))
}

func Test_isBalanced_false(t *testing.T) {
	root := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
				&TreeNode{4, nil, nil},
				&TreeNode{4, nil, nil},
			},
			&TreeNode{3, nil, nil},
		},
		&TreeNode{2, nil, nil},
	}
	assert.Equal(t, false, isBalanced(root))
}

func Test_isBalanced_nil(t *testing.T) {
	assert.Equal(t, true, isBalanced(nil))
}
