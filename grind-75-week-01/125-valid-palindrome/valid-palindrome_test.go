package valid_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome("A man, a plan, a canal: Panama"))
	assert.Equal(t, false, isPalindrome("race a car"))
	assert.Equal(t, true, isPalindrome(""))
	assert.Equal(t, false, isPalindrome("0P"))
}
