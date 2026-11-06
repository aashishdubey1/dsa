package arrays

func TwoSum(arr []int, target int) []int {
	seen := make(map[int]int)

	for i, v := range arr {
		needed := target - v
		if j, exist := seen[needed]; exist {
			return []int{j, i}
		}
		seen[v] = i
	}

	return []int{}
}
