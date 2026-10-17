package recursion

// Given an integer N, return the sum of first N natural numbers.
// Try to solve this using recursion.

func NumbersSum(n int) int {
	if n == 0 {
		return 0
	}
	return n + NumbersSum(n-1)
}
