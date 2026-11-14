package binarysearch

func RotatedSearchII(arr []int, k int) bool {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == k {
			return true
		}

		if arr[start] == arr[mid] && arr[mid] == arr[end] {
			start++
			end--
			continue
		}

		if arr[start] <= arr[mid] {
			if arr[start] <= k && k < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if arr[mid] < k && k <= arr[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}
