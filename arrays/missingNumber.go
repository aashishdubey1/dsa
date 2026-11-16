package arrays

// Given an integer array of size n containing distinct values in the range from 0 to n (inclusive),
// return the only number missing from the array within this range.

// Cycle Sort approach
// func MissingNumber(arr []int) int {
// 	i := 0
// 	for i < len(arr) {
// 		curr := arr[i]
// 		if arr[i] < len(arr) && arr[i] != arr[curr] {
// 			arr[i], arr[curr] = arr[curr], arr[i]
// 		} else {
// 			i++
// 		}
// 	}
// 	for i := range arr {
// 		if arr[i] != i {
// 			return i
// 		}
// 	}
// 	return len(arr)
// }
//

// XOR approach
// func MissingNumbers(arr []int) int {
// 	n := len(arr)
// 	xorSum := n
// 	for i := range arr {
// 		xorSum ^= i ^ arr[i]
// 	}
// 	return xorSum
// }

func MissingNumbers(arr []int) int {
	n := len(arr)
	expectedSum := (n * (n + 1) / 2)
	currSum := 0
	for _, v := range arr {
		currSum += v
	}
	return expectedSum - currSum
}
