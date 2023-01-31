package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	left := 1
	right := 1
	l := len(nums)
	result := make([]int, l)

	result[0] = 1
	for i := 1; i < l; i++ {
		left = left * nums[i-1]
		result[i] = left
	}
	for i := l - 2; i >= 0; i-- {
		right = right * nums[i+1]
		result[i] = result[i] * right
	}

	return result
}
