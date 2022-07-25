package implement_queue_using_stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Stack_1(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	assert.Equal(t, 1, obj.Peek())
	assert.Equal(t, false, obj.Empty())
	assert.Equal(t, 1, obj.Pop())
	assert.Equal(t, true, obj.Empty())
}

func Test_Stack_2(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	assert.Equal(t, 1, obj.Peek())
	assert.Equal(t, 1, obj.Pop())
	assert.Equal(t, false, obj.Empty())
}
