package main

import "testing"

func TestMajorityBit(t *testing.T) {
	readings := []string{"0100", "1100", "1010", "1000", "0011", "1001", "0001"}
	var expected map[int]byte = map[int]byte{0: '1', 1: '0', 2: '0', 3: '1'}
	for i, v := range expected {
		actual, _ := get_majority_bit(readings, i)
		if actual != v {
			t.Errorf("Testing position %v, expected %v, got %v", i, v, actual)
		}
	}
}
