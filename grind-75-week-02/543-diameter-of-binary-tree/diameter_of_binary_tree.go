package diameter_of_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	maxDepthL := maxDepth(root.Left)
	maxDepthR := maxDepth(root.Right)

	return 1 + int(math.Max(float64(maxDepthL), float64(maxDepthR)))
}
func diameterOfBinaryTree(root *TreeNode) int {
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	return l + r
}
