package sorting

func CycleSort(arr []int) []int {
	i := 0
	for i < len(arr) {
		curr := arr[i] - 1
		if arr[i] != arr[curr] {
			arr[i], arr[curr] = arr[curr], arr[i]
		} else {
			i++
		}
	}
	return arr
}
