package arrays

func SortZeroOneTwo(arr []int) []int {
	zero := 0
	two := len(arr) - 1
	i := 0

	for i < two {
		if arr[i] == 0 {
			arr[i], arr[zero] = arr[zero], arr[i]
			zero++
			i++
		} else if arr[i] == 1 {
			i++
		} else if arr[i] == 2 {
			arr[i], arr[two] = arr[two], arr[i]
			two--
		}
	}
	return arr
}
