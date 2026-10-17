package recursion

func CheckSorted(arr []int) bool {
	return isSorted(arr, 1)
}

func isSorted(arr []int, i int) bool {
	if i >= len(arr) {
		return true
	}
	if arr[i] < arr[i-1] {
		return false
	}
	return isSorted(arr, i+1)
}
