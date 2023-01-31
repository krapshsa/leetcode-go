package permutations

func dfs(nums []int, used *[]bool, answers *[][]int, answer *[]int) {
	for i := 0; i < len(nums); i++ {
		if (*used)[i] {
			continue
		}

		(*used)[i] = true
		*answer = append(*answer, nums[i])

		if len(nums) == len(*answer) {
			newAnswer := make([]int, len(*answer))
			copy(newAnswer, *answer)
			*answers = append(*answers, newAnswer)
		} else {
			dfs(nums, used, answers, answer)
		}

		(*used)[i] = false
		*answer = (*answer)[0 : len(*answer)-1]
	}
}

func permute(nums []int) [][]int {
	answers := make([][]int, 0)
	answer := make([]int, 0)
	used := make([]bool, len(nums))

	dfs(nums, &used, &answers, &answer)

	return answers
}
