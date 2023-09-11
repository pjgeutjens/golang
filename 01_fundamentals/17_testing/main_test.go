package main

import "testing"

func TestAdd(t *testing.T) {
	// arrange
	l, r, expected := 1, 2, 3

	// act
	got := Add(l, r)

	// assert
	if got != expected {
		t.Errorf("Add(%d, %d) = %d; expected %d", l, r, got, expected)
	}
}
