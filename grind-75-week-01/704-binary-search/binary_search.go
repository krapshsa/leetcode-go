package binary_search

func binarySearch(nums []int, target int, left int, right int) int {
	if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}

	m := (right + left) / 2
	if nums[m] == target {
		return m
	} else if nums[m] > target {
		return binarySearch(nums, target, left, m-1)
	} else {
		return binarySearch(nums, target, m+1, right)
	}
}

func search(nums []int, target int) int {
	l := len(nums)
	if 0 == l {
		return -1
	}

	return binarySearch(nums, target, 0, l-1)
}
