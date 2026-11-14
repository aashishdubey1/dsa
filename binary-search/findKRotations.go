package binarysearch

// Given an integer array nums of size n,
// sorted in ascending order with distinct values.
// The array has been right rotated an unknown number of times,
// between 0 and n-1 (including).
// Determine the number of rotations performed on the array.

func FindKRotations(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start < end {
		mid := start + (end-start)/2

		if arr[mid] > arr[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}
