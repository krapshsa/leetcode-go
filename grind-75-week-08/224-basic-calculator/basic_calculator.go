package basic_calculator

func calculate(s string) int {
	l := len(s)
	nums := make([]int, 0, l)
	signs := make([]int, 0, l)
	num := 0
	sign := 1
	result := 0
	for _, c := range s {
		if c-'0' >= 0 && c-'0' <= 9 {
			digit := int(c - '0')
			num = num*10 + digit
			continue
		}

		if c == ' ' {
			continue
		}

		result += sign * num
		num = 0
		if c == '+' {
			sign = 1
			continue
		} else if c == '-' {
			sign = -1
		} else if c == '(' {
			nums = append(nums, result)
			signs = append(signs, sign)
			result = 0
			sign = 1
		} else if c == ')' {
			result = nums[len(nums)-1] + signs[len(signs)-1]*result
			signs = signs[0 : len(signs)-1]
			nums = nums[0 : len(nums)-1]
		}
	}

	return result + sign*num
}
