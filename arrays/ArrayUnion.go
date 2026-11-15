package arrays

//
// Given two sorted arrays nums1 and nums2,
// return an array that contains the union of these two arrays.
// The elements in the union must be in ascending order.

func UnionArray(arr1, arr2 []int) []int {
	i := 0
	j := 0
	result := make([]int, 0, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			if len(result) == 0 || result[len(result)-1] != arr1[i] {
				result = append(result, arr1[i])
			}
			i++
		} else if arr2[j] < arr1[i] {
			if len(result) == 0 || result[len(result)-1] != arr2[j] {
				result = append(result, arr2[j])
			}
			j++
		} else {
			if len(result) == 0 || result[len(result)-1] != arr1[i] {
				result = append(result, arr1[i])
			}
			i++
			j++
		}
	}

	if i < len(arr1) {
		addElements(&result, arr1[i:])
	}

	if j < len(arr2) {
		addElements(&result, arr2[j:])
	}
	return result
}

func addElements(final *[]int, toAdd []int) {
	for _, v := range toAdd {
		if len(*final) == 0 || (*final)[len(*final)-1] != v {
			*final = append(*final, v)
		}
	}
}
