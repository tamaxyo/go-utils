package stack

type Stack struct {
	top  *element
	size int
}

type element struct {
	v    interface{}
	next *element
}

func NewStack() *Stack {
	return &Stack{
		top:  nil,
		size: 0,
	}
}

func (s *Stack) Push(v interface{}) {
	s.top = &element{v, s.top}
	s.size++
}

func (s *Stack) Pop() (v interface{}) {
	if s.size > 0 {
		v, s.top = s.top.v, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.size > 0 {
		return s.top.v
	}
	return nil
}
