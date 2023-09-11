package binarysearch

import "fmt"

func search(haystack []int, needle, left, right int) bool {
	fmt.Printf("[%d] %d %d", needle, left, right)
	if left > right {
		return false
	}

	mid := left + (right-left)/2
	if haystack[mid] == needle {
		return true
	}
	if haystack[mid] > needle {
		return search(haystack, needle, left, mid-1)
	} else {
		return search(haystack, needle, mid+1, right)
	}
}

func BinarySearch(haystack []int, needle int) bool {
	return search(haystack, needle, 0, len(haystack)-1)
}
