package arrays

// Given an integer array nums, move all the 0's to the end of the array.
// The relative order of the other elements must remain the same.
// This must be done in place, without making a copy of the array.
// Input: nums = [0, 1, 4, 0, 5, 2]
// Output: [1, 4, 5, 2, 0, 0]

// cool , first i solve this with my own thinking
// like spot zeors remove them and add to the end
// that takes lot's of effort, spoting , removing etc
// as now append doesn't create new array but for large array
// where capacity got reached it can create new array
//
// then after a hint solve using two pointer approach
// instead of moving zero move all non zero element to the front

func MoveZeros(arr []int) []int {
	count := 0
	for i, v := range arr {
		if v != 0 {
			arr[i], arr[count] = arr[count], arr[i]
			count++
		}
	}
	return arr
}

// func MoveZeros(arr []int) []int {
// 	count := 0
// 	i := 0
// 	for i < len(arr) {
// 		if arr[i] == 0 {
// 			count++
// 			arr = append(arr[:i], arr[i+1:]...)
// 		} else {
// 			i++
// 		}
// 	}
//
// 	for range count {
// 		arr = append(arr, 0)
// 	}
//
// 	return arr
// }
//
