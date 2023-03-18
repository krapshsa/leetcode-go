package find_the_divisibility

func divisibilityArray(word string, m int) []int {
	l := len(word)
	answer := make([]int, l)
	remainder := 0

	for i, r := range word {
		digit := int(r - '0')

		remainder = remainder*10 + digit
		remainder = remainder % m

		if remainder == 0 {
			answer[i] = 1
		} else {
			answer[i] = 0
		}
	}

	return answer
}
