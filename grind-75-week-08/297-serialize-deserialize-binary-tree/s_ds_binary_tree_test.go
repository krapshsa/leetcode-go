package serialize_deserialize_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SDSBinaryTreeSuite struct {
	suite.Suite
}

func TestSDSBinaryTreeTestSuite(t *testing.T) {
	suite.Run(t, new(SDSBinaryTreeSuite))
}

func (s *SDSBinaryTreeSuite) SetupTest() {
}

func (s *SDSBinaryTreeSuite) Test_1() {
	root := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil},
		},
	}
	serializer := Constructor()
	deserializer := Constructor()
	data := serializer.serialize(root)
	ans := deserializer.deserialize(data)
	assert.Equal(s.T(), ans, root, "")
}
