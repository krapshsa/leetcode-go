package lca_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(cur *TreeNode, p *TreeNode, path *[]*TreeNode) bool {
	if cur == nil {
		return false
	}

	*path = append(*path, cur)
	if cur == p {
		return true
	}

	resultLeft := dfs(cur.Left, p, path)
	if resultLeft {
		return true
	}

	resultRight := dfs(cur.Right, p, path)
	if resultRight {
		return true
	}

	*path = (*path)[0 : len(*path)-1]
	return false
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP := make([]*TreeNode, 0)
	pathQ := make([]*TreeNode, 0)
	cur := root
	dfs(cur, p, &pathP)
	dfs(cur, q, &pathQ)

	var lca *TreeNode = nil
	for i := 0; i < len(pathP) && i < len(pathQ); i++ {
		if pathP[i] == pathQ[i] {
			lca = pathP[i]
		} else {
			break
		}
	}

	return lca
}
