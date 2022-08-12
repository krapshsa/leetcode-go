package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func traversal(root *TreeNode, cur int) {
	if root == nil {
		if cur > max {
			max = cur
		}
		return
	}

	traversal(root.Left, cur+1)
	traversal(root.Right, cur+1)
}

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	max = 0
	traversal(root, 0)
	return max
}
