package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	l := len(nums)
	skip := false

	sort.Ints(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < l-1; j++ {
			if (j > i+1) && nums[j] == nums[j-1] {
				continue
			}

			for k := j + 1; k < l; k++ {
				if (k > j+1) && nums[k] == nums[k-1] {
					continue
				}

				sum := nums[i] + nums[j] + nums[k]
				if sum > 0 {
					skip = true
				}
				if sum == 0 {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
				}
			}

			if skip {
				skip = false
				continue
			}
		}
	}

	return ans
}
