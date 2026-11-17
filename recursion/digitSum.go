package recursion

func DigitSum(n int) int {
	if n == 0 {
		return n
	}

	return n%10 + DigitSum(n/10)
}
