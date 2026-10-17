package recursion

func ReverseArr(arr []int) []int {
	return reverse(arr, 0, len(arr)-1)
}

func reverse(arr []int, start, end int) []int {
	if start >= end {
		return arr
	}
	arr[start], arr[end] = arr[end], arr[start]
	return reverse(arr, start+1, end-1)
}
