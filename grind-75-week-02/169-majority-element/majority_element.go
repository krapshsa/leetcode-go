package majority_element

func majorityElement(nums []int) int {
	var maj int
	cnt := 0
	for _, num := range nums {
		if cnt == 0 {
			maj = num
			cnt++
			continue
		}

		if maj == num {
			cnt++
		} else {
			cnt--
		}
	}
	return maj
}
