package bootcamp

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v interface{}) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}
