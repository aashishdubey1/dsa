package arrays

func MaxSubArr(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]

	for _, v := range nums[1:] {
		currSum = max(v, currSum+v)
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}
