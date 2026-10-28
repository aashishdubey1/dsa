package arrays

// Given an integer array nums sorted in non-decreasing order,
// remove all duplicates in-place so that each unique element appears only once.
// Return the number of unique elements in the array.

func RemoveDuplicates(arr []int) int {
	pos := 0
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[pos] {
			arr[i], arr[pos+1] = arr[pos+1], arr[i]
			pos++
			count++
		}
	}
	return count
}
