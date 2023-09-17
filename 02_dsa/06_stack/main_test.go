package stack

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}
	_, err := s.Peek()
	if err != nil {
		t.Errorf("expect peek on empty stack to return an error")
	}
	_, err = s.Pop()
	if err != nil {
		t.Errorf("expect pop on empty stack to return an error")
	}
	s.Push(5)
	s.Push(6)
	s.Push(7)
	got, err := s.Pop()
	if err != nil || got != 7 {
		t.Errorf("expected pop to return 7, got %v", got)
	}
	got = s.Size()
	if got != 2 {
		t.Errorf("expected size to be 2, got %v", got)
	}
	got, _ = s.Peek()
	if got != 6 {
		t.Errorf("expected peeked element to be 6, got %v", got)
	}
	got = s.Size()
	if got != 2 {
		t.Errorf("expected size to be 3, got %v", got)
	}
	got = s.IsEmpty()
	if got != false {
		t.Errorf("Expected stack to have elements, isEmpty returned %v", got)
	}
	s.Pop()
	s.Pop()
	got = s.IsEmpty()
	if got != true {
		t.Errorf("Expected stack to be empty, isEmpty returned %v", got)
	}
}
