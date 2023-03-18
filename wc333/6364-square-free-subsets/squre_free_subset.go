package square_free_subsets

import (
	"math/big"
)

func squareFreeSubsets(nums []int) int {
	squareFreeNums := make([]int, 0, len(nums))
	for _, num := range nums {
		if num == 1 || big.NewInt(int64(num)).ProbablyPrime(1) {
			squareFreeNums = append(squareFreeNums, num)
		}
	}

	n := len(squareFreeNums)

	ans := 1
	for i := 0; i < n; i++ {
		ans *= 2
	}

	return ans - 1
}
