package middle_of_the_linked_list

import (
	"github.com/stretchr/testify/assert"
	. "leetcode/util"
	"testing"
)

func Test_middleNode_odd(t *testing.T) {
	node := CreateLinkedList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 3, middleNode(node).Val)
}

func Test_middleNode_even(t *testing.T) {
	node := CreateLinkedList([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, 4, middleNode(node).Val)
}
