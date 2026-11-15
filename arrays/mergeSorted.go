package arrays

func MergeTwoSortedArray(arr1, arr2 []int, m, n int) []int {

	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && arr1[i] > arr2[j] {
			arr1[k] = arr1[i]
			i--
		} else {
			arr1[k] = arr2[j]
			j--
		}
		k--
	}

	return arr1
}
