package arrays

//Given a binary array nums, return the maximum number of consecutive 1s in the array.
// A binary array is an array that contains only 0s and 1s

func FindMaxConsecutiveOnes(arr []int) int {
	max := 0
	counter := 0

	for _, v := range arr {
		if v == 1 {
			counter++
			if counter > max {
				max = counter
			}
		} else {
			counter = 0
		}
	}
	return max
}
