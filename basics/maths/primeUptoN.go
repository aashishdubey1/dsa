package maths

func PrimeUpToN(n int) int {

	count := 0
	for i := range n {
		if IsPrime(i) {
			count++
		}
	}
	return count
}
