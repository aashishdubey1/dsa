package arrays

func Reverse(arr []int) []int {
	revArr := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		revArr = append(revArr, arr[i])
	}
	return revArr
}
