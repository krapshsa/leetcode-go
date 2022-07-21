package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmpLeft := root.Left
	tmpRight := root.Right

	root.Left = invertTree(tmpRight)
	root.Right = invertTree(tmpLeft)

	return root
}
