package minimum_window_substring

type answer struct {
	len   int
	start int
	end   int
}

func minWindow(s string, t string) string {
	lenS := len(s)
	lenT := len(t)

	if lenT > lenS {
		return ""
	}

	targetRuneCnt := make(map[rune]int)
	for _, v := range t {
		targetRuneCnt[v]++
	}

	ans := answer{len: lenS + 1, start: 0, end: 0}
	subStringRuneCnt := make(map[rune]int)
	need := len(targetRuneCnt)
	have := 0
	j := -1
	for i := 0; i <= lenS-lenT; i++ {
		if i > 0 {
			getOff := rune(s[i-1])
			subStringRuneCnt[getOff]--
			if subStringRuneCnt[getOff] < targetRuneCnt[getOff] {
				have--
			}
		}

		if have == need {
			newLen := j - i + 1
			if newLen < ans.len {
				ans.start = i
				ans.end = j
				ans.len = newLen
			}
			continue
		}

		for j < lenS-1 {
			j++
			getOn := rune(s[j])
			subStringRuneCnt[getOn]++
			if subStringRuneCnt[getOn] == targetRuneCnt[getOn] {
				have++
			}
			if have == need {
				newLen := j - i + 1
				if newLen < ans.len {
					ans.start = i
					ans.end = j
					ans.len = newLen
				}
				break
			}
		}
	}

	if ans.len == lenS+1 {
		return ""
	} else {
		return s[ans.start : ans.end+1]
	}
}
