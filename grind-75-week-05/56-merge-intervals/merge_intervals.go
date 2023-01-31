package merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	answer := make([][]int, 0)
	answer = append(answer, intervals[0])

	for i := 1; i < len(intervals); i++ {
		preIntervalIndex := len(answer) - 1
		if answer[preIntervalIndex][1] >= intervals[i][0] {
			if answer[preIntervalIndex][1] < intervals[i][1] {
				answer[preIntervalIndex][1] = intervals[i][1]
			}
		} else {
			answer = append(answer, intervals[i])
		}
	}

	return answer
}
