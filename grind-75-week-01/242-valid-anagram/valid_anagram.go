package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runeMap := make(map[rune]int)
	for _, value := range s {
		runeMap[value]++
	}

	for _, value := range t {
		cnt, ok := runeMap[value]
		if ok && cnt > 0 {
			runeMap[value]--
		} else {
			return false
		}
	}

	return true
}
