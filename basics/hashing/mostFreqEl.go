package hashing

func MostFrequentElement(arr []int) int {

	freqMap := make(map[int]int)

	for _, v := range arr {
		freqMap[v]++
	}

	mostFreq := 0
	mostFreqEl := 0

	for k, v := range freqMap {
		if v > mostFreq {
			mostFreq = v
			mostFreqEl = k
		} else if v == mostFreq && k < mostFreqEl {
			mostFreqEl = k
		}
	}

	return mostFreqEl
}
