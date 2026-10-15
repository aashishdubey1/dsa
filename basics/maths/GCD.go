package maths

func GCD(n1, n2 int) int {
	smallNum := min(n1, n2)
	gcd := 0
	for i := 1; i <= smallNum; i++ {
		cd := n1%i == 0 && n2%i == 0
		if cd && i > gcd {
			gcd = i
		}
	}
	return gcd
}

// this solution works but for large numbers it's not the optimized one

func GCD2(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
