package binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, curHeight int, values *[]int) {
	if nil == node {
		return
	}

	if len(*values) <= curHeight {
		*values = append(*values, node.Val)
	} else {
		(*values)[curHeight] = node.Val
	}

	dfs(node.Left, curHeight+1, values)
	dfs(node.Right, curHeight+1, values)
}

func rightSideView(root *TreeNode) []int {
	rightNodeValue := make([]int, 0)

	dfs(root, 0, &rightNodeValue)

	return rightNodeValue
}
