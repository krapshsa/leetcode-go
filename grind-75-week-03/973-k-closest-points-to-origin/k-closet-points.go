package k_closest_points_to_origin

type MinHeap [][]int

func (h *MinHeap) distance(i int) int {
	return (*h)[i][0]*(*h)[i][0] + (*h)[i][1]*(*h)[i][1]
}

func (h *MinHeap) swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) push(n []int) {
	*h = append(*h, n)

	// reorder
	cur := len(*h) - 1
	parent := (cur - 1) / 2
	for cur != parent {
		if h.distance(cur) > h.distance(parent) {
			break
		}

		h.swap(parent, cur)
		cur = parent
		parent = (cur - 1) / 2
	}
}

func (h *MinHeap) extract() []int {
	n := (*h)[0]
	last := len(*h) - 1
	(*h)[0] = (*h)[last]
	*h = (*h)[0:last]

	// reorder from top
	h.reorderFromTop()

	return n
}

func (h *MinHeap) reorderFromTop() {
	last := len(*h) - 1
	cur := 0
	for cur < last {
		leftChild := cur*2 + 1
		rightChild := cur*2 + 2

		minIdx := cur
		if leftChild <= last && h.distance(leftChild) < h.distance(minIdx) {
			minIdx = leftChild
		}
		if rightChild <= last && h.distance(rightChild) < h.distance(minIdx) {
			minIdx = rightChild
		}

		if minIdx == cur {
			break
		}

		h.swap(minIdx, cur)
		cur = minIdx
	}
}

func kClosest(points [][]int, k int) [][]int {
	minHeap := MinHeap{}

	for i := 0; i < len(points); i++ {
		minHeap.push(points[i])
	}

	var result [][]int
	for i := 0; i < k; i++ {
		n := minHeap.extract()
		result = append(result, n)
	}
	return result
}
