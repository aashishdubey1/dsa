package sorting

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i <= n-2; i++ {
		j := i + 1
		for j > 0 {
			if arr[j] >= arr[j-1] {
				break
			}
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	return arr
}
