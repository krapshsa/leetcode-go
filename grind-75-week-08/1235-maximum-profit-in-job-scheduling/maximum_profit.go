package maximum_profit_in_job_scheduling

import "sort"

type Schedule struct {
	startTime int
	endTime   int
	profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	l := len(startTime)
	schedules := make([]Schedule, l)

	maxEndTime := 0
	for i := 0; i < l; i++ {
		schedules[i] = Schedule{
			startTime: startTime[i],
			endTime:   endTime[i],
			profit:    profit[i],
		}

		if maxEndTime < endTime[i] {
			maxEndTime = endTime[i]
		}
	}

	sort.Slice(schedules, func(i, j int) bool {
		return schedules[i].endTime < schedules[j].endTime
	})

	dp := make([]int, l)
	curMax := 0
	for newIndex, schedule := range schedules {
		// find old end < new start
		oldIndex := newIndex - 1
		for ; oldIndex >= 0; oldIndex-- {
			if schedules[oldIndex].endTime <= schedule.startTime {
				break
			}
		}

		var profit int
		if oldIndex >= 0 {
			profit = dp[oldIndex] + schedule.profit
		} else {
			//dp[newIndex] = schedule.profit
			profit = schedule.profit
		}

		if profit > curMax {
			curMax = profit
		}

		dp[newIndex] = curMax
	}

	return dp[l-1]
}
