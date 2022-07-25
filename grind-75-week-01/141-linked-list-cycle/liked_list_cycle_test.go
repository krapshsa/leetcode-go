package linked_list_cycle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasCycle_normal(t *testing.T) {
	n1 := &ListNode{3, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{0, nil}
	n4 := &ListNode{4, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	assert.Equal(t, true, hasCycle(n1))
}

func Test_hasCycle_two_node_loop(t *testing.T) {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n1.Next = n2
	n2.Next = n1
	assert.Equal(t, true, hasCycle(n1))
}

func Test_hasCycle_single_node(t *testing.T) {
	head := &ListNode{1, nil}
	assert.Equal(t, false, hasCycle(head))
}

func Test_hasCycle_two_node_no_loop(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}
	assert.Equal(t, false, hasCycle(head))
}
