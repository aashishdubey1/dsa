package sorting

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(first, second []int) []int {
	mix := make([]int, 0, len(first)+len(second))

	i, j := 0, 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			mix = append(mix, first[i])
			i++
		} else {
			mix = append(mix, second[j])
			j++
		}
	}
	mix = append(mix, first[i:]...)
	mix = append(mix, second[j:]...)

	return mix
}
