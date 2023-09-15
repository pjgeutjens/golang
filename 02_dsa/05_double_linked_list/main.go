package doublelinkedlist

import (
	"fmt"
	"strconv"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (l LinkedList) Length() int {
	return l.len
}

func (l LinkedList) Print() string {
	result := "[ "
	n := l.head
	for n != nil {
		result += strconv.Itoa(n.value)
		result += ", "
		n = n.next
	}
	result += " ]"

	return result
}

func (l *LinkedList) Add(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head, l.tail = node, node
		l.len++
		return
	}
	node.prev = l.tail
	l.tail.next = node
	l.tail = node
	l.len++
}

func (l LinkedList) Get(idx int) (int, error) {
	n := l.head
	for i := 0; i < idx; i++ {
		if n.next == nil {
			return 0, fmt.Errorf("index out of bounds")
		}
		n = n.next
	}
	return n.value, nil
}

func (l LinkedList) Contains(value int) bool {
	n := l.head
	for n != nil {
		if n.value == value {
			return true
		}
		n = n.next
	}
	return false
}

func (l *LinkedList) InsertAt(idx int, value int) error {
	node := &Node{value: value}
	if idx == 0 {
		node.next = l.head
		l.head = node
		l.len++
		return nil
	}

	n := l.head
	for i := 0; i < idx; i++ {
		if n.next == nil {
			return fmt.Errorf("idx out of bounds")
		}
		n = n.next
	}

	node.prev = n.prev
	n.prev.next = node
	node.next = n
	n.prev = node
	l.len++
	return nil
}
