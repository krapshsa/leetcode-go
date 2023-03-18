package minimum_impossible_or

import "sort"

func minImpossibleOR(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	max := nums[l-1]
	final := 1
	for max > 0 {
		max = max / 2
		final *= 2
	}

	for i := 1; i < max; i++ {

	}

	return max
}
