package left_and_right_sum_diff

func leftRigthDifference(nums []int) []int {
	l := len(nums)

	leftSum := make([]int, l)
	rightSum := make([]int, l)
	answer := make([]int, l)

	for i := 1; i < l; i++ {
		leftSum[i] = leftSum[i-1] + nums[i-1]
	}

	for j := l - 2; j >= 0; j-- {
		rightSum[j] = rightSum[j+1] + nums[j+1]
	}

	for i := 0; i < l; i++ {
		answer[i] = leftSum[i] - rightSum[i]
		if answer[i] < 1 {
			answer[i] = answer[i] * -1
		}
	}

	return answer
}
