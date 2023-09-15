package doublelinkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := LinkedList{}
	// testing Length
	l := list.Length()
	if l != 0 {
		t.Errorf("Expect list length to be 0, got %d", list.Length())
	}

	// testing Add
	list.Add(1)
	fmt.Println(list.Print())
	list.Add(2)
	fmt.Println(list.Print())
	list.Add(3)
	fmt.Println(list.Print())
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
	if val, err := list.Get(3); err == nil || val != 0 {
		t.Errorf("Expected error, got nil")
	}

	// test Contains
	if val := list.Contains(2); val != true {
		t.Errorf("Expected Contains(2) to return true, got %v", val)
	}
	if val := list.Contains(4); val != false {
		t.Errorf("Expected Contains(4) to return false, got %v", val)
	}
	// test InsertAt
	err := list.InsertAt(1, 15)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val, err := list.Get(1); err != nil || val != 15 {
		t.Errorf("Expected GetAt(1) to return 15, got %v", val)
	}
	if val, err := list.Get(2); err != nil || val != 2 {
		t.Errorf("Expected GetAt(2) to return 2, got %v", val)
	}
	l = list.Length()
	if l != 4 {
		t.Errorf("Expect list length to be 4, got %d", list.Length())
	}
	err = list.InsertAt(0, 9)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val, err := list.Get(0); err != nil || val != 9 {
		t.Errorf("Expected GetAt(1) to return 15, got %v", val)
	}
	// Test InsertAt error
	err = list.InsertAt(10, 5)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	//	.Run("RemoveAt", func(t *testing.T) {
	//			// Test
	//			val, err := list.RemoveAt(1)
	//			if err != nil {
	//				t.Errorf("Expected no error, got %v", err)
	//			}
	//			if val != 2 {
	//				t.Errorf("Expected RemoveAt(1) to return 2, got %v", val)
	//			}
	//			if val, err := list.Get(1); err != nil || val != 3 {
	//				t.Errorf("Expected Get(2) to return 3, got %v", val)
	//			}
	//		})
	//
	//	t.Run("InsertAt", func(t *testing.T) {
	//		// Test InsertAt
	//		err := list.InsertAt(1, 4)
	//		if err != nil {
	//			t.Errorf("Expected no error, got %v", err)
	//		}
	//		if val, err := list.Get(1); err != nil || val != 4 {
	//			t.Errorf("Expected Get(2) to return 4, got %v", val)
	//		}
	//		if val, err := list.Get(2); err != nil || val != 3 {
	//			t.Errorf("Expected Get(1) to return 3, got %v", val)
	//		}
	//	})
	//
	// // Test RemoveAt error
	// _, err = list.RemoveAt(5)
	//
	//	if err == nil {
	//		t.Errorf("Expected error, got nil")
	//	}
	//
	// Test GetAt error
	//
	//	if _, err := list.Get(5); err == nil {
	//		t.Errorf("Expected error, got nil")
	//	}
}
