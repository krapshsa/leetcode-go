package maximum_difference

func minMaxDifference(num int) int {
	digits := make([]int, 0)

	n := num
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	l := len(digits)
	diff := make([]int, l)
	max := make([]int, l)
	min := make([]int, l)
	copy(max, digits)
	copy(min, digits)
	target := -1
	for i := l - 1; i >= 0; i-- {
		if 9 == digits[i] {
			continue
		} else {
			if target == -1 {
				target = digits[i]
			}

			if target == digits[i] {
				max[i] = 9
			}
		}
	}

	target = -1
	for i := l - 1; i >= 0; i-- {
		if 0 == digits[i] {
			continue
		} else {
			if target == -1 {
				target = digits[i]
			}
			if target == digits[i] {
				min[i] = 0
			}
		}
	}

	for i := 0; i < l; i++ {
		diff[i] = max[i] - min[i]
	}

	ans := 0
	mul := 1
	for i := 0; i < l; i++ {
		ans += diff[i] * mul
		mul *= 10
	}

	return ans
}
