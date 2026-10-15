package maths

func Divisors(n int) []int {
	divisors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
