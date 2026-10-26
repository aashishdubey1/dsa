package arrays

// Given an integer array nums, rotate the array to the left by one.
// Note: There is no need to return anything, just modify the given array.
// Input: nums = [1, 2, 3, 4, 5]
// Output: [2, 3, 4, 5, 1]

func RotateArrayByOne(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	first := arr[0]
	for i := 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[n-1] = first

	return arr
}
