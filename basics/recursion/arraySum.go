package recursion

func ArraySum(arr []int) int {
	return arrSum(arr, 0)
}

func arrSum(arr []int, i int) int {
	if i >= len(arr) {
		return 0
	}
	return arr[i] + arrSum(arr, i+1)
}
