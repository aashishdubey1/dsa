package recursion

func SearchAll(nums []int, target int) []int {

	result := []int{}

	return searchAll(nums, result, target, 0)
}

func searchAll(nums, result []int, target, i int) []int {

	if i >= len(nums) {
		return result
	}

	if nums[i] == target {
		result = append(result, i)
	}

	return searchAll(nums, result, target, i+1)

}
