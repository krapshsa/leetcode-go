package reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseList_long(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	expected := &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	assert.Equal(t, expected, reverseList(list))
}

func Test_reverseList_short(t *testing.T) {
	list := &ListNode{1, &ListNode{2, nil}}
	expected := &ListNode{2, &ListNode{1, nil}}
	assert.Equal(t, expected, reverseList(list))
}

func Test_reverseList_nil(t *testing.T) {
	assert.Equal(t, nil, reverseList(nil))
}
