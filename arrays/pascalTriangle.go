package arrays

func factorial(num int) int {
	result := 1
	for i := 2; i <= num; i++ {
		result *= i
	}
	return result
}

func PascalTriangle(r, c int) int {

	n := r - 1
	k := c - 1

	return factorial(n) / (factorial(k) * factorial(n-k))
}
