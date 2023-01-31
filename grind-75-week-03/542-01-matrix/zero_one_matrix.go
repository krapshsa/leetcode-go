package zero_one_matrix

func updateMatrix(mat [][]int) [][]int {
	height := len(mat)
	width := len(mat[0])

	// init
	result := make([][]int, height)
	for i := range result {
		result[i] = make([]int, width)
	}

	q := make([][2]int, 0, width*height)

	// put all 0 into q
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if mat[i][j] == 0 {
				result[i][j] = 0
				q = append(q, [2]int{i, j})
			} else {
				result[i][j] = 9999
			}
		}
	}

	// pop & update
	for len(q) != 0 {
		i := q[0][0]
		j := q[0][1]
		q = q[1:]

		distance := result[i][j] + 1
		if j-1 >= 0 && result[i][j-1] == 9999 {
			result[i][j-1] = distance
			q = append(q, [2]int{i, j - 1})
		}
		if j+1 < width && result[i][j+1] == 9999 {
			result[i][j+1] = distance
			q = append(q, [2]int{i, j + 1})
		}
		if i-1 >= 0 && result[i-1][j] == 9999 {
			result[i-1][j] = distance
			q = append(q, [2]int{i - 1, j})
		}
		if i+1 < height && result[i+1][j] == 9999 {
			result[i+1][j] = distance
			q = append(q, [2]int{i + 1, j})
		}
	}

	return result
}
