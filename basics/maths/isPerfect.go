package maths

// Perfect only if
// proper divisors (excluding the number itself) add up to the number itself.

func IsPerfect(n int) bool {
	properDivisiorSum := 0
	for i := 1; i <= n/2; i++ {
		properDivisiorSum += i
	}
	return properDivisiorSum == n
}
