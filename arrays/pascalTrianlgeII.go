package arrays

//Given an integer r, return all the values in the rth row (1-indexed) in Pascal's Triangle in correct order.

func PascalTriangleII(r int) []int {
	n := r - 1
	result := []int{}

	for k := range r {
		val := factorial(n) / (factorial(k) * factorial(n-k))
		result = append(result, val)
		k++
	}

	return result
}
