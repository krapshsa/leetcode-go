package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list1RestNode := list1
	list2RestNode := list2
	dummy := &ListNode{0, nil}
	var tail *ListNode = dummy
	var candidate **ListNode = nil

	for list1RestNode != nil || list2RestNode != nil {
		if list1RestNode == nil {
			candidate = &list2RestNode
		} else if list2RestNode == nil {
			candidate = &list1RestNode
		} else {
			if list1RestNode.Val > list2RestNode.Val {
				candidate = &list2RestNode
			} else {
				candidate = &list1RestNode
			}
		}

		tail.Next = &ListNode{(*candidate).Val, nil}
		*candidate = (*candidate).Next
		tail = tail.Next
	}

	return dummy.Next
}
