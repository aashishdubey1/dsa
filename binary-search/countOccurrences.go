package binarysearch

func CountOccurrences(arr []int, target int) int {
	return upperBound(arr, target) - lowerBound(arr, target)
}
