package arrays

// Given two sorted arrays, nums1 and nums2,
// return an array containing the intersection of these two arrays.
// Each element in the result must appear as many times as it appears in both arrays;
// that is, if an element appears x times in nums1 and y times in nums2,
// it should appear min(x, y) times in the result.

func IntersectionArray(arr1, arr2 []int) []int {
	i := 0
	j := 0
	result := make([]int, 0, min(len(arr1), len(arr2)))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			result = append(result, arr1[i])
			i++
			j++
		} else {
			if arr1[i] < arr2[j] {
				i++
			} else {
				j++
			}
		}
	}
	return result
}
