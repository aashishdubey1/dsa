package maths

func Factorial(n int) int {
	ans := 1
	for n > 0 {
		ans *= n
		n--
	}
	return ans
}
