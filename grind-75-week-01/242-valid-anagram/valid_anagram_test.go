package valid_anagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	assert.Equal(t, true, isAnagram("anagram", "nagaram"))
	assert.Equal(t, false, isAnagram("rat", "car"))
	assert.Equal(t, false, isAnagram("aabbb", "aaabb"))
}
