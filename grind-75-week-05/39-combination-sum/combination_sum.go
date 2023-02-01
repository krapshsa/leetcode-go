package combination_sum

import "sort"

func sum(candidates []int, target int, answer *[][]int, stack *[]int) {
	for i, v := range candidates {
		if v < target {
			*stack = append(*stack, v)
			sum(candidates[i:], target-v, answer, stack)
			*stack = (*stack)[0 : len(*stack)-1]
		} else if v == target {
			*stack = append(*stack, v)
			result := make([]int, len(*stack))
			copy(result, *stack)
			*answer = append(*answer, result)
			*stack = (*stack)[0 : len(*stack)-1]
			return
		} else {
			break
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	answer := make([][]int, 0)
	stack := make([]int, 0)

	sort.Ints(candidates)
	sum(candidates, target, &answer, &stack)

	return answer
}
