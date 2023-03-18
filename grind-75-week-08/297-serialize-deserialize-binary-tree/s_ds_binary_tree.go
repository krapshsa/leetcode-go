package serialize_deserialize_binary_tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) dfsSerialize(node *TreeNode, fields *[]string) {
	if node == nil {
		*fields = append(*fields, "n")
		return
	}

	*fields = append(*fields, strconv.Itoa((*node).Val))
	this.dfsSerialize(node.Left, fields)
	this.dfsSerialize(node.Right, fields)
}

func (this *Codec) dfsDeserialize(node **TreeNode, fields []string, cur *int) {
	val := fields[*cur]
	if val == "n" {
		return
	}

	intVal, _ := strconv.Atoi(val)
	*node = &TreeNode{intVal, nil, nil}

	*cur = *cur + 1
	this.dfsDeserialize(&(**node).Left, fields, cur)

	*cur = *cur + 1
	this.dfsDeserialize(&(**node).Right, fields, cur)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	fields := make([]string, 0)
	this.dfsSerialize(root, &fields)

	return strings.Join(fields, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	var root *TreeNode
	fields := strings.Fields(data)
	cur := 0
	this.dfsDeserialize(&root, fields, &cur)

	return root
}
