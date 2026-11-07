package binarysearch

//Given a sorted array nums and an integer x.
//Find the floor and ceil of x in nums.
//The floor of x is the largest element in the array which is smaller than or equal to x.
//The ceiling of x is the smallest element in the array greater than or equal to x.
//If no floor or ceil exists, output -1.

func FloorAndCeil(arr []int, target int) []int {

	start := 0
	end := len(arr) - 1

	floor := -1
	ceil := -1

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == target {
			return []int{arr[mid], arr[mid]}
		}
		if arr[mid] < target {
			floor = arr[mid]
			start = mid + 1
		} else {
			ceil = arr[mid]
			end = mid - 1
		}
	}
	return []int{floor, ceil}
}
