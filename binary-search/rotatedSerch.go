package binarysearch

func RotatedSearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid
		}

		if arr[start] <= arr[mid] {
			if arr[start] <= target && target < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if arr[mid] < target && target <= arr[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
