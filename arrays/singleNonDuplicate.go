package arrays

func SingleNonDuplicate(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := start + (end-start)/2

		if mid%2 == 1 {
			mid--
		}

		if nums[mid] == nums[mid+1] {
			start = mid + 2
		} else {
			end = mid
		}
	}
	return nums[start]
}
