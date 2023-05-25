package samples

// binarySearchIndex returns the index at which elem can be inserted
// in the sorted array a, such that a remains sorted.
func BinarySearchIndex(a []int, elem int) int {
	left, right := 0, len(a)
	for left < right {
		mid := (left + right) / 2
		if a[mid] < elem {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// User implement binary search find the position of an element that can be inserted in an array
// using a Binary search and return the index
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// If the target is not found, return the position where it can be inserted
	return low
}
