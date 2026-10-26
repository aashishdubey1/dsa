package arrays

// Given an integer array of size n containing distinct values in the range from 0 to n (inclusive),
// return the only number missing from the array within this range.

func MissingNumber(arr []int) int {
	i := 0
	for i < len(arr) {
		curr := arr[i]
		if arr[i] < len(arr) && arr[i] != arr[curr] {
			arr[i], arr[curr] = arr[curr], arr[i]
		} else {
			i++
		}
	}
	for i := range arr {
		if arr[i] != i {
			return i
		}
	}
	return len(arr)
}
