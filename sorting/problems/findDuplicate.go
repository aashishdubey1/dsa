package problems

// Given an array of integers nums containing n + 1 integers
// where each integer is in the range [1,n] inclusive
// There is only one repeted number in nums, return this repeted number
//
// must solve the problem without modifying the array nums and usees only constant extra space

func FindDuplicate(arr []int) int {
	i := 0
	for i < len(arr) {
		curr := arr[i] - 1
		if arr[i] != arr[curr] {
			arr[i], arr[curr] = arr[curr], arr[i]
		} else {
			i++
		}
	}
	for i := range arr {
		if arr[i] != i+1 {
			return arr[i]
		}
	}
	return -1
}
