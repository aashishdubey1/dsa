package maths

func CountDigit(n int) int {
	if n == 0 {
		return 1
	}
	var count int
	for n > 0 {
		count++
		n /= 10
	}
	return count
}
