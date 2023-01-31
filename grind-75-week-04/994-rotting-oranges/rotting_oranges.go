package rotting_oranges

import "math"

func orangesRotting(grid [][]int) int {
	length := len(grid)
	width := len(grid[0])
	q := make([][2]int, 0, length*width)
	time := make([][]int, length)
	max := 0

	for i := 0; i < length; i++ {
		time[i] = make([]int, width)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 1 {
				time[i][j] = math.MaxInt
			}
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			}
		}
	}

	for len(q) > 0 {
		rottingOrangeX, rottingOrangeY := q[0][0], q[0][1]
		q = q[1:]
		for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			orangeX := rottingOrangeX + v[0]
			orangeY := rottingOrangeY + v[1]
			if orangeX >= 0 && orangeX < length && orangeY >= 0 && orangeY < width {
				if grid[orangeX][orangeY] == 1 && time[orangeX][orangeY] == math.MaxInt {
					time[orangeX][orangeY] = time[rottingOrangeX][rottingOrangeY] + 1
					q = append(q, [2]int{orangeX, orangeY})
				}
			}
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if time[i][j] > max {
				max = time[i][j]
			}
		}
	}

	if max == math.MaxInt {
		return -1
	} else {
		return max
	}
}
