package stack

import "fmt"

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Size() == 0 {
		return nil, fmt.Errorf("empty stack")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

func (s Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack) Size() int {
	return len(s.elements)
}
