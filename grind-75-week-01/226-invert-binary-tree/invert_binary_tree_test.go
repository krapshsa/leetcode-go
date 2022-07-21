package invert_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func createBinaryTree(node []int, shift int, level int) *TreeNode {
	idx := shift + int(math.Pow(2, float64(level))) - 1
	if idx >= len(node) {
		return nil
	}

	return &TreeNode{
		node[idx],
		createBinaryTree(node, 2*shift, level+1),
		createBinaryTree(node, 2*shift+1, level+1),
	}
}

func Test_invertTree(t *testing.T) {
	expect := createBinaryTree([]int{4, 2, 7, 1, 3, 6, 9}, 0, 0)
	root := createBinaryTree([]int{4, 7, 2, 9, 6, 3, 1}, 0, 0)

	assert.Equal(t, expect, invertTree(root))
}
