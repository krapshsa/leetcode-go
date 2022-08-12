package middle_of_the_linked_list

import (
	. "leetcode/util"
)

func middleNode(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	slow := dummy
	fast := dummy
	for fast != nil {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
		slow = slow.Next
	}
	return slow
}
