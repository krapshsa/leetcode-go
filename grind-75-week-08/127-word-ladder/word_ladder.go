package word_ladder

func canTransform(sourceWord string, targetWord string) bool {
	diff := 0
	for i := 0; i < len(sourceWord); i++ {
		if sourceWord[i] != targetWord[i] {
			diff++
		}

		if diff == 2 {
			return false
		}
	}

	return true
}
func ladderLength(beginWord string, endWord string, wordList []string) int {
	l := len(wordList)
	times := make(map[string]int, l+1)

	endInList := false
	for i := 0; i < l; i++ {
		if wordList[i] == endWord {
			endInList = true
			break
		}
	}
	if !endInList {
		return 0
	}

	queue := make([]string, 0, l+1)
	queue = append(queue, beginWord)
	times[beginWord] = 1
	cur := 0
	for len(queue) > cur {
		curWord := queue[cur]
		cur++
		curLevel := times[curWord]

		for i := 0; i < l; i++ {
			word := wordList[i]
			if times[wordList[i]] > 0 {
				continue
			}

			if canTransform(curWord, word) {
				if wordList[i] == endWord {
					return curLevel + 1
				}

				// BFS
				queue = append(queue, word)
				times[word] = curLevel + 1
			}
		}
	}

	return 0
}
