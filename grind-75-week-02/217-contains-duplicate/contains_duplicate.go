package contains_duplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			return true
		} else {
			m[num] = 1
		}
	}

	return false
}
