package recursion

func FindElement(arr []int, target int) int {
	return findEl(arr, target, 0)
}

func findEl(arr []int, target, i int) int {
	if i >= len(arr) {
		return -1
	}
	if arr[i] == target {
		return i
	}

	return findEl(arr, target, i+1)
}
