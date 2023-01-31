package min_stack

type MinStack struct {
	node []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	min := val
	if len((*this).node) > 0 {
		pMin := (*this).GetMin()
		if pMin < min {
			min = pMin
		}
	}
	(*this).node = append((*this).node, val)
	(*this).min = append((*this).min, min)
}

func (this *MinStack) Pop() {
	newLen := len((*this).node) - 1
	(*this).node = (*this).node[:newLen]
	(*this).min = (*this).min[:newLen]
}

func (this *MinStack) Top() int {
	return (*this).node[len((*this).node)-1]
}

func (this *MinStack) GetMin() int {
	return (*this).min[len((*this).min)-1]
}
