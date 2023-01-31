package coin_change

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	l := len(coins)
	sort.Ints(coins)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < l; j++ {
			coin := coins[j]
			if i-coin >= 0 {
				result := 1 + dp[i-coin]
				if dp[i] > result {
					dp[i] = result
				}
			} else {
				break
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[amount]
	}
}
