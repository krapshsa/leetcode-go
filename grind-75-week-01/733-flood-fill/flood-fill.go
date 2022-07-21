package flood_fill

type FIFO struct {
	content [][2]int
}

func (fifo *FIFO) Enqueue(idx [2]int) {
	fifo.content = append(fifo.content, idx)
}

func (fifo *FIFO) Dequeue() [2]int {
	ret := fifo.content[0]
	fifo.content = fifo.content[1:]
	return ret
}

func (fifo *FIFO) Len() int {
	return len(fifo.content)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	fifo := &FIFO{make([][2]int, 0, 50)}
	maxY := len(image) - 1
	maxX := len(image[0]) - 1

	fifo.Enqueue([2]int{sr, sc})
	for fifo.Len() > 0 {
		idx := fifo.Dequeue()
		y := idx[0]
		x := idx[1]

		if y > maxY || y < 0 || x < 0 || x > maxX {
			continue
		}
		if image[y][x] != originalColor {
			continue
		}
		if image[y][x] == color {
			continue
		}

		image[y][x] = color

		fifo.Enqueue([2]int{y + 1, x})
		fifo.Enqueue([2]int{y - 1, x})
		fifo.Enqueue([2]int{y, x + 1})
		fifo.Enqueue([2]int{y, x - 1})
	}

	return image
}
