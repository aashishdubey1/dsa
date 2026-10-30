package arrays

func RearragenArray(nums []int) []int {
	pos := 0
	neg := 1

	res := make([]int, len(nums))

	for _, v := range nums {
		if v > 0 {
			res[pos] = v
			pos += 2
		} else if v < 0 {
			res[neg] = v
			neg += 2

		}
	}
	return res
}
