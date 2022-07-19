package valid_palindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	l := len(s)

	if 0 == l || 1 == l {
		return true
	}

	head := 0
	tail := l - 1

	for head < tail {
		headRune := rune(s[head])
		if !unicode.IsLetter(headRune) && !unicode.IsDigit(headRune) {
			head++
			continue
		}

		tailRune := rune(s[tail])
		if !unicode.IsLetter(tailRune) && !unicode.IsDigit(tailRune) {
			tail--
			continue
		}

		if unicode.ToLower(headRune) != unicode.ToLower(tailRune) {
			break
		}

		head++
		tail--
	}

	return head >= tail
}
