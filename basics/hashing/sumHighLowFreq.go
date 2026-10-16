package hashing

func SumHighestAndLowestFrequency(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	freqMap := make(map[int]int)

	for _, v := range nums {
		freqMap[v]++
	}

	highFreq := 0
	lowFreq := len(nums)

	for _, v := range freqMap {
		if v > highFreq {
			highFreq = v
		}
		if v < lowFreq {
			lowFreq = v
		}
	}

	return highFreq + lowFreq
}
