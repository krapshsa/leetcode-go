package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(nums []int) *ListNode {
	dummy := ListNode{0, nil}
	now := &dummy
	for _, num := range nums {
		now.Next = &ListNode{num, nil}
		now = now.Next
	}
	return dummy.Next
}
