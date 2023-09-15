package linked_list

import "testing"

func TestLinkedList(t *testing.T) {
	list := LinkedList{}
	l := list.Length()
	if l != 0 {
		t.Errorf("Expect list length to be 0, got %d", list.Length())
	}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	l = list.Length()
	if l != 3 {
		t.Errorf("Expect list length to be 3, got %d", list.Length())
	}
	// Test Add and Get
	if val, err := list.Get(0); err != nil || val != 1 {
		t.Errorf("Expected Get(0) to return 1, got %v", val)
	}
	if val, err := list.Get(1); err != nil || val != 2 {
		t.Errorf("Expected Get(1) to return 2, got %v", val)
	}
	if val, err := list.Get(2); err != nil || val != 3 {
		t.Errorf("Expected Get(2) to return 3, got %v", val)
	}
	// Test
	val, err := list.RemoveAt(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val != 2 {
		t.Errorf("Expected RemoveAt(1) to return 2, got %v", val)
	}
	if val, err := list.Get(1); err != nil || val != 3 {
		t.Errorf("Expected Get(1) to return 3, got %v", val)
	}
	// Test InsertAt
	err = list.InsertAt(1, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val, err := list.Get(1); err != nil || val != 4 {
		t.Errorf("Expected Get(2) to return 4, got %v", val)
	}
	if val, err := list.Get(2); err != nil || val != 3 {
		t.Errorf("Expected Get(1) to return 3, got %v", val)
	}
	// Test RemoveAt error
	_, err = list.RemoveAt(5)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	// Test GetAt error
	if _, err := list.Get(5); err == nil {
		t.Errorf("Expected error, got nil")
	}
}
