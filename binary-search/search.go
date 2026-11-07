package binarysearch

// search on sorted array

func Search(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if target == arr[mid] {
			return mid
		}
		if target > arr[mid] {
			start = mid + 1
		} else if target < arr[mid] {
			end = mid - 1
		}
	}
	return -1
}

// func search(arr []int, target int) int {
// 	start := 0
// 	end := len(arr) - 1
//
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if target == arr[mid] {
// 			return mid
// 		}
// 		if target > arr[mid] {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
// }
