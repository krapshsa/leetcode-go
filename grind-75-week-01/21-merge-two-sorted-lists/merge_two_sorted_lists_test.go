package merge_two_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	expected := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}
	assert.Equal(t, expected, mergeTwoLists(list1, list2))
}

func Test_mergeTwoLists_empty(t *testing.T) {
	expected := (*ListNode)(nil)
	list1 := (*ListNode)(nil)
	list2 := (*ListNode)(nil)
	assert.Equal(t, expected, mergeTwoLists(list1, list2))
}

func Test_mergeTwoLists_single_empty(t *testing.T) {
	list1 := (*ListNode)(nil)
	list2 := &ListNode{0, nil}
	expected := &ListNode{0, nil}
	assert.Equal(t, expected, mergeTwoLists(list1, list2))
}
