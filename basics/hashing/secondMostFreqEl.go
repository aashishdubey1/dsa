package hashing

func SecondMostFrequentElement(nums []int) int {

	freqMap := make(map[int]int)

	for _, v := range nums {
		freqMap[v]++
	}

	mostFreq := 0

	for _, v := range freqMap {
		if v > mostFreq {
			mostFreq = v
		}
	}

	secondMostFreq := 0
	secondMostFreqEl := -1

	for k, v := range freqMap {
		if v > secondMostFreq && v < mostFreq {
			secondMostFreq = v
			secondMostFreqEl = k
		} else if v == secondMostFreq && v < mostFreq && k < secondMostFreqEl {
			secondMostFreqEl = k
		}
	}
	return secondMostFreqEl
}
