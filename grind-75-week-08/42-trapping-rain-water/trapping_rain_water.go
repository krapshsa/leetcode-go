package trapping_rain_water

func trap(height []int) int {
	maxIndex, _ := findMax(height)
	cnt := len(height)

	trapped := 0
	left := 0
	for i := 1; i < maxIndex; i++ {
		if height[i] >= height[left] {
			left = i
		} else {
			trapped += height[left] - height[i]
		}
	}

	right := cnt - 1
	for i := cnt - 2; i > maxIndex; i-- {
		if height[i] >= height[right] {
			right = i
		} else {
			trapped += height[right] - height[i]
		}
	}

	return trapped
}

func findMax(height []int) (int, int) {
	idx := 0
	max := 0
	for i, v := range height {
		if v > max {
			max = v
			idx = i
		}
	}
	return idx, max
}
