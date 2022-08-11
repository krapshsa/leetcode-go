package longest_palindrome

func longestPalindrome(s string) int {
	runesCnt := make(map[rune]int)
	stringLen := len(s)
	result := 0

	for i := 0; i < stringLen; i++ {
		c := rune(s[i])
		if runesCnt[c] > 0 {
			result += 2
			runesCnt[c] = 0
		} else {
			runesCnt[c]++
		}
	}

	if result < stringLen {
		return result + 1
	} else {
		return result
	}
}
