package find_all_anagrams

func isAnagram(m map[rune]int) bool {
	for _, value := range m {
		if value != 0 {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	lenP := len(p)
	lenS := len(s)
	result := make([]int, 0)
	runeMap := make(map[rune]int)

	if lenP > lenS {
		return result
	}

	for _, r := range p {
		if _, ok := runeMap[r]; ok {
			runeMap[r]++
		} else {
			runeMap[r] = 1
		}
	}

	for i := 0; i < lenP; i++ {
		r := rune(s[i])
		if _, ok := runeMap[r]; !ok {
			continue
		}

		runeMap[r]--
	}

	if isAnagram(runeMap) {
		result = append(result, 0)
	}

	for i := 1; i <= lenS-lenP; i++ {
		previousRune := rune(s[i-1])
		if _, ok := runeMap[previousRune]; ok {
			runeMap[previousRune]++
		}

		currentRune := rune(s[i+lenP-1])
		if _, ok := runeMap[currentRune]; !ok {
			continue
		}

		runeMap[currentRune]--

		if isAnagram(runeMap) {
			result = append(result, i)
		}
	}

	return result
}
