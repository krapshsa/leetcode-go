package merge_two_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	. "leetcode/util"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list1 := CreateLinkedList([]int{1, 2, 4})
	list2 := CreateLinkedList([]int{1, 3, 4})
	expected := CreateLinkedList([]int{1, 1, 2, 3, 4, 4})
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
	list2 := CreateLinkedList([]int{0})
	expected := CreateLinkedList([]int{0})
	assert.Equal(t, expected, mergeTwoLists(list1, list2))
}
