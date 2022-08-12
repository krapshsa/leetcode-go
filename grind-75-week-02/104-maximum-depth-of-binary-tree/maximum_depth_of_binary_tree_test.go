package maximum_depth_of_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Queue struct {
	content []*TreeNode
}

func (q *Queue) push(node *TreeNode) {
	q.content = append(q.content, node)
}

func (q *Queue) pop() *TreeNode {
	var ret *TreeNode
	ret, q.content = q.content[0], q.content[1:]
	return ret
}

func newQueue() *Queue {
	return &Queue{make([]*TreeNode, 0, 10000)}
}

func createBinaryTree(nums []int) *TreeNode {
	q := newQueue()
	var num int
	var head *TreeNode

	num, nums = nums[0], nums[1:]
	if num == -1 {
		head = nil
	} else {
		head = &TreeNode{num, nil, nil}
	}
	q.push(head)

	for len(nums) > 0 {
		node := q.pop()
		num, nums = nums[0], nums[1:]
		var left *TreeNode
		if num == -1 {
			left = nil
		} else {
			left = &TreeNode{num, nil, nil}
		}
		node.Left = left
		q.push(left)

		if len(nums) > 0 {
			num, nums = nums[0], nums[1:]
			var right *TreeNode
			if num == -1 {
				right = nil
			} else {
				right = &TreeNode{num, nil, nil}
			}
			node.Right = right
			q.push(right)
		}
	}

	return head
}

func Test_maxDepth(t *testing.T) {
	assert.Equal(t, 3, maxDepth(createBinaryTree([]int{3, 9, 20, -1, -1, 15, 7})))
	assert.Equal(t, 2, maxDepth(createBinaryTree([]int{1, -1, 2})))
}
