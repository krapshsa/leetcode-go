package merge_k_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MergeSuite struct {
	suite.Suite
}

func TestMergeTestSuite(t *testing.T) {
	suite.Run(t, new(MergeSuite))
}

func (s *MergeSuite) SetupTest() {
}

func (s *MergeSuite) Test_1() {
	lists := []*ListNode{
		&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		&ListNode{2, &ListNode{6, nil}},
	}

	mergedNodes := mergeKLists(lists)
	actual := make([]int, 0)
	for mergedNodes != nil {
		actual = append(actual, mergedNodes.Val)
		mergedNodes = mergedNodes.Next
	}

	assert.Equal(s.T(), []int{1, 1, 2, 3, 4, 4, 5, 6}, actual, "")
}

func (s *MergeSuite) Test_2() {
	var expected *ListNode = nil
	assert.Equal(s.T(), expected, mergeKLists([]*ListNode{}), "")
}

func (s *MergeSuite) Test_3() {
	var expected *ListNode = nil
	assert.Equal(s.T(), expected, mergeKLists([]*ListNode{nil}), "")
}
