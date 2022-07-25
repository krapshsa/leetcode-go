package LCA_of_a_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lowest := root

	for true {
		if lowest.Val > p.Val && lowest.Val > q.Val {
			lowest = lowest.Left
		} else if lowest.Val < p.Val && lowest.Val < q.Val {
			lowest = lowest.Right
		} else {
			break
		}
	}

	return lowest
}
