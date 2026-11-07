package binarysearch

func SearchRange(arr []int, target int) []int {
	first := lowerBound(arr, target)
	if first >= len(arr) || arr[first] != target {
		return []int{-1, -1}
	}

	last := upperBound(arr, target)

	return []int{first, last - 1}
}

func lowerBound(arr []int, target int) int {
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

func upperBound(arr []int, target int) int {
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
