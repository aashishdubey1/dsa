package maths

func CountOddDigit(n int) int {
	var count int
	for n > 0 {
		digit := n % 10
		if digit%2 != 0 {
			count++
		}
		n /= 10
	}
	return count
}
