package find_median_from_data_stream

import "container/heap"

type IntMinHeap []int
type IntMaxHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}
func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntMinHeap) Top() int {
	return h[0]
}
func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntMinHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func (h IntMaxHeap) Len() int {
	return len(h)
}
func (h IntMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntMaxHeap) Top() int {
	return h[0]
}
func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntMaxHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

type MedianFinder struct {
	minHeap *IntMinHeap
	maxHeap *IntMaxHeap
}

func Constructor() MedianFinder {
	minHeap := &IntMinHeap{}
	maxHeap := &IntMaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)

	return MedianFinder{minHeap: minHeap, maxHeap: maxHeap}
}

func (this *MedianFinder) AddNum(num int) {
	minHeapLen := this.minHeap.Len()
	maxHeapLen := this.maxHeap.Len()

	if len(*this.maxHeap) > len(*this.minHeap) {
		// add to min heap
		if this.maxHeap.Top() > num {
			n := heap.Pop(this.maxHeap)
			heap.Push(this.minHeap, n)
			heap.Push(this.maxHeap, num)
		} else {
			heap.Push(this.minHeap, num)
		}
	} else { // min length == max length
		if minHeapLen == 0 && maxHeapLen == 0 {
			// init - add to max heap
			heap.Push(this.maxHeap, num)
		} else {
			// add to max heap
			if this.minHeap.Top() < num {
				n := heap.Pop(this.minHeap)
				heap.Push(this.maxHeap, n)
				heap.Push(this.minHeap, num)
			} else {
				heap.Push(this.maxHeap, num)
			}
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	maxHeapLen := this.maxHeap.Len()
	minHeapLen := this.minHeap.Len()

	if maxHeapLen > minHeapLen {
		return float64(this.maxHeap.Top())
	}
	if 0 == maxHeapLen {
		return 0
	}
	sum := float64(this.maxHeap.Top()) + float64(this.minHeap.Top())
	return sum / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
