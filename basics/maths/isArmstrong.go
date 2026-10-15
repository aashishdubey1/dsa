package maths

func IsArmstrong(n int) bool {

	if n < 0 {
		return false
	}

	len := 0
	sum := 0
	originalNumber := n
	temp := n

	if temp == 0 {
		len = 1
	}

	for temp > 0 {
		len++
		temp /= 10
	}

	for n > 0 {
		digit := n % 10
		sum += power(digit, len)
		n /= 10
	}
	return originalNumber == sum
}

func power(base, exp int) int {
	result := 1
	for range exp {
		result *= base
	}
	return result
}
