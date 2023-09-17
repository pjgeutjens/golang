package queue

import "fmt"

type Queue struct {
	items []interface{}
}

func (s *Queue) Enqueue(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Queue) Dequeue() (interface{}, error) {
	if s.Size() == 0 {
		return nil, fmt.Errorf("empty stack")
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item, nil
}

func (s Queue) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("empty")
	}
	return s.items[0], nil
}

func (s Queue) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Queue) Size() int {
	return len(s.items)
}
