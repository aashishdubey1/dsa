package binarysearch

// Given a sorted array of nums consisting of distinct integers and a target value
// return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.

func SearchInsert(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	ans := len(arr)

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] >= target {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
