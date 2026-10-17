package recursion

func AddDigits(a int) int {
	if a <= 0 {
		return 0
	}
	return a%10 + AddDigits(a/10)
}
