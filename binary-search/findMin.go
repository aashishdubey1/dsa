package binarysearch

// Given an integer array nums of size N,
// sorted in ascending order with distinct values,
// and then rotated an unknown number of times (between 1 and N),
// find the minimum element in the array.

func FindMin(arr []int) int {
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
	return arr[start]
}
