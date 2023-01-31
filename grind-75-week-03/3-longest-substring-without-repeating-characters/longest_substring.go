package ongest_substring

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	max := 0
	left, right := 0, 0

	for ; right < len(s); right++ {
		m[s[right]]++
		for ; m[s[right]] > 1 && left < right; left++ {
			m[s[left]]--
		}
		if max < right-left+1 {
			max = right - left + 1
		}
	}

	return max
}
