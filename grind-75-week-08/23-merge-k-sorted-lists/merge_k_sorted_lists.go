package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	cur := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = &ListNode{list1.Val, nil}
			list1 = list1.Next
		} else {
			cur.Next = &ListNode{list2.Val, nil}
			list2 = list2.Next
		}
		cur = cur.Next
	}

	for list1 != nil {
		cur.Next = &ListNode{list1.Val, nil}
		list1 = list1.Next
		cur = cur.Next
	}

	for list2 != nil {
		cur.Next = &ListNode{list2.Val, nil}
		list2 = list2.Next
		cur = cur.Next
	}

	return dummyHead.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}

	oldList := make([]*ListNode, l)
	copy(oldList, lists)

	for len(oldList) > 1 {
		newList := make([]*ListNode, 0)
		for i := 0; i < len(oldList); i += 2 {
			if i+1 < len(oldList) {
				newList = append(newList, mergeTwoList(oldList[i], oldList[i+1]))
			} else {
				newList = append(newList, oldList[i])
			}
		}
		oldList = newList
	}

	return oldList[0]
}
