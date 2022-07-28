package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	var charCnt [26]int
	for _, char := range magazine {
		charCnt[char-'a']++
	}

	for _, char := range ransomNote {
		index := char - 'a'
		if charCnt[index] == 0 {
			return false
		} else {
			charCnt[index]--
		}
	}

	return true
}
