package recursion

func BubbleSort(nums []int) []int {

	return sort(nums, len(nums)-1, 0)
}

func sort(nums []int, r, c int) []int {
	if r == 0 {
		return nums
	}

	if c < r {
		if nums[c] > nums[c+1] {
			nums[c], nums[c+1] = nums[c+1], nums[c]
		}
		return sort(nums, r, c+1)
	} else {
		return sort(nums, r-1, 0)
	}
}
