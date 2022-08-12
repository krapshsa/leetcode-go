package reverse_linked_list

import . "leetcode/util"

func reverseList(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	var pre *ListNode
	pre = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
