package recursion

func Search(nums []int, target int) int {
	return search(nums, target, 0)
}

func search(nums []int, target, i int) int {
	if i >= len(nums) {
		return -1
	}
	if nums[i] == target {
		return i
	}
	return search(nums, target, i+1)
}
