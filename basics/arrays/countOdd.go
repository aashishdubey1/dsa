package arrays

func CountOdd(arr []int) int {
	var count int
	for _, v := range arr {
		if v%2 != 0 {
			count++
		}
	}
	return count
}
