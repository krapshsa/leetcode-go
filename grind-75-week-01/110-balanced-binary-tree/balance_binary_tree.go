package balanced_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func depth(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 1
	}

	balanced, lDepth := depth(node.Left)
	if !balanced {
		return false, 0
	}

	balanced, rDepth := depth(node.Right)
	if !balanced {
		return false, 0
	}

	if math.Abs(float64(lDepth-rDepth)) > 1 {
		return false, 0
	}
	return true, 1 + int(math.Max(float64(lDepth), float64(rDepth)))
}

func isBalanced(root *TreeNode) bool {
	balanced, _ := depth(root)
	return balanced
}
