package letter_combinations

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	answers := []string{""}
	lettersMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	for _, digit := range digits {
		letters := lettersMap[byte(digit)]
		newAnswers := make([]string, 0)
		for _, letter := range letters {
			for _, answer := range answers {
				newAnswer := answer + string(letter)
				newAnswers = append(newAnswers, newAnswer)
			}
		}
		answers = newAnswers
	}

	return answers
}
