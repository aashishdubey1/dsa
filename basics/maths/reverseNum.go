package maths

func ReverseNum(n int) int {
	rev := 0
	for n > 0 {
		digit := n % 10
		rev = rev*10 + digit
		n /= 10
	}
	return rev
}
