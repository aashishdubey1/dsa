package arrays

func FindMissingRepeatingNumbers(arr []int) []int {
	freqMap := make(map[int]int)
	result := make([]int, 0, 2)
	for _, v := range arr {
		freqMap[v]++
	}

	for k, v := range freqMap {
		if v == 2 {
			result = append(result, k)
			break
		}
	}

	for i := 1; i <= len(arr); i++ {
		_, exist := freqMap[i]
		if !exist {
			result = append(result, i)
			break
		}
	}
	return result
}
