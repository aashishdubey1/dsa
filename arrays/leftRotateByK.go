package arrays

// Given an integer array nums and a non-negative integer k,
// rotate the array to the left by k steps.

func RotateArray(arr []int, k int) []int {
	n := len(arr)

	if n == 0 {
		return arr
	}

	k %= n

	for range k {
		first := arr[0]
		for i := range n - 1 {
			arr[i] = arr[i+1]
		}
		arr[n-1] = first
	}
	return arr
}
