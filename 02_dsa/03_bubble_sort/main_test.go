package bubblesort

import (
	"reflect"
	"testing"
)

// write a test for bubblesort
func TestBubbleSort(t *testing.T) {
	data := []int{4, 19, 101, 7, 88, 55, 90, 12, 55}
	expected := []int{4, 7, 12, 19, 55, 55, 88, 90, 101}
	result := BubbleSort(data)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("BubbleSort(%d) = %d", data, result)
	}
}
