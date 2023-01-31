package number_of_islands

func numIslands(grid [][]byte) int {
	length := len(grid)
	width := len(grid[0])
	label := make([][]int, length)
	queue := make([][2]int, 0, length*width)
	max := 0

	for i := 0; i < length; i++ {
		label[i] = make([]int, width)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == byte('0') || label[i][j] != 0 {
				continue
			}

			max++
			label[i][j] = max
			queue = append(queue, [2]int{i, j})
			for len(queue) > 0 {
				oi, oj := queue[0][0], queue[0][1]
				queue = queue[1:len(queue)]
				direction := [4][2]int{{oi - 1, oj}, {oi + 1, oj}, {oi, oj - 1}, {oi, oj + 1}}
				for k := 0; k < 4; k++ {
					ni, nj := direction[k][0], direction[k][1]
					if ni >= 0 && nj >= 0 && ni < length && nj < width && label[ni][nj] == 0 && grid[ni][nj] == byte('1') {
						label[ni][nj] = label[oi][oj]
						queue = append(queue, [2]int{ni, nj})
					}
				}
			}
		}
	}

	return max
}
