package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(root *TreeNode, values *[]int) {
	leftNode := (*root).Left
	if leftNode != nil {
		traversal(leftNode, values)
	}

	*values = append(*values, (*root).Val)

	rightNode := (*root).Right
	if rightNode != nil {
		traversal(rightNode, values)
	}
}

func isValidBST(root *TreeNode) bool {
	inOrderValues := make([]int, 0)
	traversal(root, &inOrderValues)
	for i := 1; i < len(inOrderValues); i++ {
		if inOrderValues[i-1] >= inOrderValues[i] {
			return false
		}
	}

	return true
}
