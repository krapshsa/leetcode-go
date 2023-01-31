package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] { // left range
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right range
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}
