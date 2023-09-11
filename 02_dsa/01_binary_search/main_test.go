package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	haystack := []int{1, 3, 4, 69, 71, 90, 94, 112, 420, 1337, 69420}
	test_cases := map[int]bool{
		1:     true,
		3:     true,
		69:    true,
		81:    false,
		111:   false,
		420:   true,
		69420: true,
		69429: false,
	}

	for needle, expected := range test_cases {
		result := BinarySearch(haystack, needle)
		if result != expected {
			t.Error("Expected ", expected, " got ", result, "searching for ", needle)
		}
	}

	for needle, expected := range test_cases {
		result := RBinarySearch(haystack, needle)
		if result != expected {
			t.Error("Expected ", expected, " got ", result, "searching for ", needle)
		}
	}
}
