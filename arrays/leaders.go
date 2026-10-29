package arrays

func Leaders(arr []int) []int {

	max := arr[len(arr)-1]
	maxArr := []int{}
	maxArr = append(maxArr, max)
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > max {
			max = arr[i]
			maxArr = append(maxArr, max)
		}
	}
	n := len(maxArr) - 1
	for i := range n {
		maxArr[i], maxArr[n-i] = maxArr[n-i], maxArr[i]
	}
	return maxArr
}
