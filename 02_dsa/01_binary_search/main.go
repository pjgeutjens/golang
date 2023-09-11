package binarysearch

import "fmt"

func recurse_search(haystack []int, needle, left, right int) bool {
	fmt.Printf("[%d] %d %d", needle, left, right)
	if left > right {
		return false
	}

	mid := left + (right-left)/2
	if haystack[mid] == needle {
		return true
	}
	if haystack[mid] > needle {
		return recurse_search(haystack, needle, left, mid-1)
	} else {
		return recurse_search(haystack, needle, mid+1, right)
	}
}

func loop_search(haystack []int, needle, left, right int) bool {
	for left <= right {
		mid := left + (right-left)/2
		if haystack[mid] == needle {
			return true
		}

		if haystack[mid] > needle {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func BinarySearch(haystack []int, needle int) bool {
	return loop_search(haystack, needle, 0, len(haystack)-1)
}

func RBinarySearch(haystack []int, needle int) bool {
	return recurse_search(haystack, needle, 0, len(haystack)-1)
}
