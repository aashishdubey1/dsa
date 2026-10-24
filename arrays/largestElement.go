package arrays

func LargetElement(arr []int) int {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
