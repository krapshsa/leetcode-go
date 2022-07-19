package valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValid(t *testing.T) {
	assert.Equal(t, true, isValid("()"), "")
	assert.Equal(t, true, isValid("()[]{}"), "")
	assert.Equal(t, false, isValid("(]"), "")
}
