package bootcamp

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(v interface{}) {
	q.items = append(q.items, v)
}

func (q *Queue) Pop() interface{} {
	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}
