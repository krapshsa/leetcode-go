package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9), "9 exists in nums and its index is 4")
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2), "2 does not exist in nums so return -1")
	assert.Equal(t, 0, search([]int{5}, 5), "5 exists in nums and its index is 0")
}
