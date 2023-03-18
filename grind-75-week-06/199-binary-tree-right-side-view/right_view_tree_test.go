package binary_tree_right_side_view

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RightViewSuite struct {
	suite.Suite
}

func TestRightViewTestSuite(t *testing.T) {
	suite.Run(t, new(RightViewSuite))
}

func (s *RightViewSuite) SetupTest() {
}

func (s *RightViewSuite) Test_1() {
	root := &TreeNode{
		1,
		&TreeNode{2, nil, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil, &TreeNode{4, nil, nil}},
	}
	assert.Equal(s.T(), []int{1, 3, 4}, rightSideView(root), "")
}

func (s *RightViewSuite) Test_2() {
	root := &TreeNode{1, nil, &TreeNode{3, nil, nil}}
	assert.Equal(s.T(), []int{1, 3}, rightSideView(root), "")
}

func (s *RightViewSuite) Test_3() {
	var root *TreeNode = nil
	assert.Equal(s.T(), []int{}, rightSideView(root), "")
}
