package recursion

import "math"

func PalindromeNumber(n int) bool {
	if n < 0 {
		return false
	}
	return pHelper(n)
}

func pHelper(n int) bool {

	if n < 10 {
		return true
	}

	first := firstDigit(n)
	last := lastDigit(n)

	if first != last {
		return false
	}

	divisor := int(math.Log10(float64(n)))
	remaining := n % divisor / 10

	return pHelper(remaining)
}

func firstDigit(n int) int {

	digits := int(math.Log10(float64(n))) + 1

	divisor := int(math.Pow(10, float64(digits-1)))

	return n / divisor
}

func lastDigit(n int) int {
	return n % 10
}
