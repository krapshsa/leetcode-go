package reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	. "leetcode/util"
	"testing"
)

func Test_reverseList_long(t *testing.T) {
	list := CreateLinkedList([]int{1, 2, 3, 4, 5})
	expected := CreateLinkedList([]int{5, 4, 3, 2, 1})
	assert.Equal(t, expected, reverseList(list))
}

func Test_reverseList_short(t *testing.T) {
	list := CreateLinkedList([]int{1, 2})
	expected := CreateLinkedList([]int{2, 1})
	assert.Equal(t, expected, reverseList(list))
}

func Test_reverseList_nil(t *testing.T) {
	assert.Equal(t, (*ListNode)(nil), reverseList(nil))
}
