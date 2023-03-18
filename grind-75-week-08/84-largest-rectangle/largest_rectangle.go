package largest_rectangle

type Record struct {
	height int
	index  int
}

type Stack struct {
	records []Record
}

func (s *Stack) len() int {
	return len(s.records)
}

func (s *Stack) push(r Record) {
	s.records = append(s.records, r)
}

func (s *Stack) pop() {
	s.records = s.records[:s.len()-1]
}

func (s *Stack) top() Record {
	r := s.records[s.len()-1]
	return r
}

func largestRectangleArea(heights []int) int {
	l := len(heights)
	stack := Stack{make([]Record, 0, l)}
	max := 0

	// iterate all heights
	for i := 0; i < l; i++ {
		height := heights[i]
		newIndex := i
		for stack.len() > 0 {
			lastRecord := stack.top()
			if lastRecord.height <= height {
				break
			}
			stack.pop()
			newIndex = lastRecord.index
			area := (i - lastRecord.index) * lastRecord.height
			if area > max {
				max = area
			}
		}
		stack.records = append(stack.records, Record{height: height, index: newIndex})
	}

	for stack.len() > 0 {
		lastRecord := stack.top()
		stack.pop()
		area := (l - lastRecord.index) * lastRecord.height
		if area > max {
			max = area
		}
	}

	return max
}
