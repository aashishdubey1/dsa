package arrays

import "math"

func SecondLargestElement(arr []int) int {
	largest := LargetElement(arr)
	secondLargest := 0
	for _, v := range arr {
		if v > secondLargest && v < largest {
			secondLargest = v
		}
	}
	return secondLargest
}

// Another approach
func SecondLargestElement2(arr []int) int {

	if len(arr) < 2 {
		return -1
	}

	first := math.MinInt64
	second := math.MinInt64

	for _, v := range arr {
		if v > first {
			second = first
			first = v
		} else if v > second && v != first {
			second = v
		}
	}

	if second == math.MinInt64 {
		return -1
	}

	return second
}
