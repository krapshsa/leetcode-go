package partition_equal_subset_sum

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)

	if nums[0] == target {
		return true
	}
	if nums[0] < target {
		dp[nums[0]] = true
	}
	l := len(nums)
	for i := 1; i < l; i++ {
		num := nums[i]
		for j := target; j >= 0 && j >= num; j-- {
			if j == num {
				dp[j] = true
			} else if j > num {
				if dp[j-num] == true {
					dp[j] = true
				}
			}
		}
	}

	return dp[target]
}
