package linked_list

import "fmt"

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
}

func (list LinkedList) Length() int {
	n, i := list.head, 0
	for n != nil {
		i++
		n = n.next
	}
	return i
}

func (list *LinkedList) Add(value int) {
	node := &Node{value: value}
	if list.head == nil {
		list.head = node
		return
	} else {
		n := list.head
		for n.next != nil {
			n = n.next
		}
		n.next = node
	}
}

func (list *LinkedList) RemoveAt(idx int) (int, error) {
	// start at the head
	previous := list.head
	if previous == nil {
		return 0, fmt.Errorf("list is empty")
	}
	if idx == 0 {
		list.head = list.head.next
		return previous.value, nil
	}
	// find the node previous to where I need to RemoveAt
	for i := 0; i < idx-1; i++ {
		if previous.next == nil {
			return 0, fmt.Errorf("idx out of bounds")
		}
		previous = previous.next
	}
	// find the target node
	n := previous.next
	// get the target node value
	// set the previous node next value to the target node next value
	previous.next = n.next
	return n.value, nil
}

func (list *LinkedList) InsertAt(idx int, value int) error {
	node := &Node{value: value}
	l := list.Length()
	if idx > l {
		return fmt.Errorf("idx out of bounds")
	}
	n := list.head
	if idx == 0 {
		node.next = list.head
		list.head = node
		return nil
	}

	for i := 0; i < idx-1; i++ {
		if n.next == nil {
			return fmt.Errorf("idx out of bounds")
		}
		n = n.next
	}

	node.next = n.next
	n.next = node
	return nil
}

func (list LinkedList) Get(idx int) (int, error) {
	if idx > list.Length()-1 {
		return 0, fmt.Errorf("index out of range")
	}
	n := list.head
	for i := 0; n.next != nil && i < idx; i++ {
		n = n.next
	}
	return n.value, nil
}
