package linked_list_cycle

import . "leetcode/util"

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
