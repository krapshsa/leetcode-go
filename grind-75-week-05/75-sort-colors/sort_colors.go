package sort_colors

func sortColors(nums []int) {
	colorMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		color := nums[i]
		cnt, ok := colorMap[color]
		if ok {
			colorMap[color] = cnt + 1
		} else {
			colorMap[color] = 1
		}
	}

	// loop all colors
	idx := 0
	for _, color := range []int{0, 1, 2} {
		for i := 0; i < colorMap[color]; i++ {
			nums[idx] = color
			idx++
		}
	}
}
