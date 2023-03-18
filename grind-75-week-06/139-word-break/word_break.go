package word_break

func wordBreak(s string, wordDict []string) bool {
	sLen := len(s)
	dp := make([]bool, sLen+1)
	dp[sLen] = true

	for i := sLen - 1; i >= 0; i-- {
		for _, word := range wordDict {
			wLen := len(word)

			// longer than string
			if wLen+i > sLen {
				continue
			}

			// cannot find matched left substring
			if dp[i+wLen] == false {
				continue
			}

			subString := s[i : i+wLen]
			if word == subString {
				dp[i] = true
			}
		}
	}

	return dp[0]
}
