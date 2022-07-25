package implement_queue_using_stacks

type Stack struct {
	content []int
}

func (s *Stack) Push(x int) {
	s.content = append(s.content, x)
}

func (s *Stack) Pop() int {
	len := len(s.content)
	x := s.content[len-1]
	s.content = s.content[:len-1]
	return x
}

func (s *Stack) Peek() int {
	return s.content[len(s.content)-1]
}

func (s *Stack) Empty() bool {
	return len(s.content) == 0
}

type MyQueue struct {
	normalStack   *Stack
	reversedStack *Stack
}

func Constructor() MyQueue {
	q := MyQueue{
		&Stack{make([]int, 0, 100)},
		&Stack{make([]int, 0, 100)},
	}
	return q
}

func (q *MyQueue) check() {
	if 0 == len(q.reversedStack.content) {
		l := len(q.normalStack.content)
		for i := 0; i < l; i++ {
			q.reversedStack.Push(q.normalStack.Pop())
		}
	}
}

func (q *MyQueue) Push(x int) {
	q.normalStack.Push(x)
}

func (q *MyQueue) Pop() int {
	q.check()
	return q.reversedStack.Pop()
}

func (q *MyQueue) Peek() int {
	q.check()
	return q.reversedStack.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.normalStack.Empty() && q.reversedStack.Empty()
}
