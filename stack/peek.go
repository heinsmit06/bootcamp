package bootcamp

type Stack2 struct {
	items []interface{}
}

func NewStack2() *Stack {
	return &Stack{}
}

func (s *Stack) Push2(v interface{}) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop2() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func (s *Stack) Len2() int {
	return len(s.items)
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	return s.items[len(s.items)-1]
}
