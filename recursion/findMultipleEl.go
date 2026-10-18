package recursion

func FindMultipleElement(arr []int, target int) []int {
	return findMlEl(arr, []int{}, target, 0)
}

func findMlEl(arr, ans []int, target, i int) []int {
	if i >= len(arr) {
		return ans
	}
	if arr[i] == target {
		ans = append(ans, i)
	}

	return findMlEl(arr, ans, target, i+1)
}
