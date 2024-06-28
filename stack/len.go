package bootcamp

type Stack1 struct {
	items []interface{}
}

func NewStack1() *Stack {
	return &Stack{}
}

func (s *Stack) Push1(v interface{}) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop1() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func (s *Stack) Len() int {
	return len(s.items)
}
