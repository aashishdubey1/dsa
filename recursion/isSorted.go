package recursion

func IsSorted(nums []int) bool {
	return accending(nums) || decending(nums)
}

func accending(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	first := nums[0]
	next := nums[1]
	return first <= next && accending(nums[1:])
}

func decending(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	first := nums[0]
	next := nums[1]
	return first >= next && decending(nums[1:])
}
