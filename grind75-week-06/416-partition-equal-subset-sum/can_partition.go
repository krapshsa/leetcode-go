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
	dp := make([][]bool, len(nums))
	for i, _ := range dp {
		dp[i] = make([]bool, target+1)
	}

	dp[0][0] = true
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		curSum += nums[i]
		bound := target
		if curSum < target {
			bound = curSum
		}
		for j := 0; j <= bound; j++ {
			if dp[i-1][j] == true {
				dp[i][j] = true
				continue
			}
			if nums[i] == j {
				dp[i][j] = true
				continue
			}
			if nums[i] < j {
				if dp[i-1][j-nums[i]] == true {
					dp[i][j] = true
				}
			}
		}
	}

	return dp[len(nums)-1][target]
}
