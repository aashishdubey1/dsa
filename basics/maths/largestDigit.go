package maths

func LargestDigit(n int) int {
	max := 0
	for n > 0 {
		digit := n % 10
		if digit > max {
			max = digit
		}
		n /= 10
	}
	return max
}
