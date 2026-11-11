package binarysearch

func UpperBound(arr []int, target int) int {

	start := 0
	end := len(arr) - 1

	ans := len(arr)

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] > target {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
