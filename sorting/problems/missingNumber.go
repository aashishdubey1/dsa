package problems

// Given an array containing n distinct numbers in the range [0,n]
// return the only number in the range that is missing from the arr
// could you implement using only O(n) and O(1) space

func MissingNumber(arr []int) int {
	i := 0
	for i < len(arr) {
		correct := arr[i]
		if arr[i] < len(arr) && arr[i] != arr[correct] {
			arr[i], arr[correct] = arr[correct], arr[i]
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
