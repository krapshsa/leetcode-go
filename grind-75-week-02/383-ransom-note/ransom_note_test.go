package ransom_note

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	assert.Equal(t, false, canConstruct("a", "b"))
	assert.Equal(t, false, canConstruct("aa", "ab"))
	assert.Equal(t, true, canConstruct("aa", "aab"))
}
