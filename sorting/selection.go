package sorting

func SelectionSort(arr []int) []int {
	for i := range arr {
		lastIndex := len(arr) - 1 - i
		maxIndex := findMax(arr, 0, lastIndex)
		arr[maxIndex], arr[lastIndex] = arr[lastIndex], arr[maxIndex]
	}
	return arr
}

func findMax(arr []int, start, last int) int {
	max := start
	for i := start; i <= last; i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	return max
}
