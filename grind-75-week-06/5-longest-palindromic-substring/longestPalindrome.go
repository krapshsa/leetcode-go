package longest_palindromic_substring

func longestPalindrome(s string) string {
	l := len(s)
	dp := make([][]bool, l)
	for i, _ := range dp {
		dp[i] = make([]bool, l)
	}

	maxLen := 1
	maxStart := 0
	maxEnd := 0
	for end := 0; end < l; end++ {
		for start := end; start >= 0; start-- {
			if s[start] != s[end] {
				continue
			}

			if start == end {
				dp[start][end] = true
			} else if end-start == 1 {
				dp[start][end] = true
			} else {
				if dp[start+1][end-1] == true {
					dp[start][end] = true
				} else {
					continue
				}
			}

			if end-start+1 > maxLen {
				maxLen = end - start + 1
				maxStart = start
				maxEnd = end
			}
		}
	}

	return s[maxStart : maxEnd+1]
}
