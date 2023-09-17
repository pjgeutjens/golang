package queue

import "testing"

func TestQueue(t *testing.T) {
	s := Queue{}
	_, err := s.Peek()
	if err == nil {
		t.Errorf("expect peek on empty queue to return an error")
	}
	_, err = s.Dequeue()
	if err == nil {
		t.Errorf("expect pop on empty queue to return an error")
	}
	s.Enqueue(5)
	s.Enqueue(6)
	s.Enqueue(7)
	got, err := s.Dequeue()
	if err != nil || got != 5 {
		t.Errorf("expected pop to return 5, got %v", got)
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
		t.Errorf("Expected queue to have elements, isEmpty returned %v", got)
	}
	s.Dequeue()
	s.Dequeue()
	got = s.IsEmpty()
	if got != true {
		t.Errorf("Expected queue to be empty, isEmpty returned %v", got)
	}
}
