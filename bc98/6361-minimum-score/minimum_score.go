package minimum_score

import "sort"

func minimizeSum(nums []int) int {
	l := len(nums)
	sort.Ints(nums)

	if l == 3 {
		return 0
	}

	high := nums[l-1] - nums[0]
	if nums[l-2]-nums[1] < high {
		high = nums[l-2] - nums[1]
	}
	if nums[l-1]-nums[2] < high {
		high = nums[l-1] - nums[2]
	}
	if nums[l-3]-nums[0] < high {
		high = nums[l-3] - nums[0]
	}

	return high
}
