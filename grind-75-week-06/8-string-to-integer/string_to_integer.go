package string_to_integer

func isDigit(b byte) bool {
	diff := getDigit(b)
	if diff >= 0 && diff <= 9 {
		return true
	} else {
		return false
	}
}

func getDigit(b byte) int {
	return int(b - byte('0'))
}

func isOut(s string, min bool) bool {
	if len(s) > 10 {
		return true
	}
	if len(s) < 10 {
		return false
	}

	maxDigits := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
	if min {
		maxDigits[9] = 8
	}
	for i := 0; i < 10; i++ {
		digit := getDigit(s[i])
		if digit > maxDigits[i] {
			return true
		} else if digit < maxDigits[i] {
			return false
		}
	}

	return false
}

func myAtoi(s string) int {
	mul := 1
	answer := 0

	start := -1
	end := -1
	for i := 0; i < len(s); i++ {
		b := s[i]
		if isDigit(b) {
			start = i
			end = i
			break
		}
	}

	if start == -1 {
		return 0
	}
	if start > 0 {
		for i := 0; i < start-1; i++ {
			if byte(' ') != s[i] {
				return 0
			}
		}

		if s[start-1] == byte('-') {
			mul = -1
		} else if s[start-1] == byte('+') || s[start-1] == byte(' ') {
			mul = 1
		} else {
			return 0
		}
	}

	for i := start + 1; i < len(s); i++ {
		b := s[i]
		if isDigit(b) {
			end = i
		} else {
			break
		}
	}

	newStart := start
	for i := start; i <= end; i++ {
		if s[i] == byte('0') {
			newStart++
		} else {
			break
		}
	}
	start = newStart

	if start > end {
		return 0
	}

	if isOut(s[start:end+1], mul == -1) {
		if mul == 1 {
			return 2147483647
		} else {
			return -2147483648
		}
	}

	for i := start; i <= end; i++ {
		answer = answer*10 + getDigit(s[i])
	}

	return answer * mul
}
