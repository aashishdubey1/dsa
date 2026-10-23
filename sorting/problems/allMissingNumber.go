package problems

func AllMissingNumbers(arr []int) []int {
	i := 0
	for i < len(arr) {
		curr := arr[i] - 1
		if arr[i] != arr[curr] {
			arr[i], arr[curr] = arr[curr], arr[i]
		} else {
			i++
		}
	}

	missingNumbers := []int{}

	for i := range arr {
		if arr[i] != i+1 {
			missingNumbers = append(missingNumbers, i+1)
		}
	}

	return missingNumbers
}
