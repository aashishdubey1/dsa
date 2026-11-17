package recursion

import (
	"fmt"
	"math"
)

func ReverseNumber(n int) int {
	digit := int(math.Log10(float64(n))) + 1
	fmt.Println("digit", digit)
	return helper(n, digit)
}

func helper(n, digit int) int {
	if n%10 == n {
		return n
	}
	rem := n % 10
	return rem*int(math.Pow(10, float64(digit-1))) + helper(n/10, digit-1)
}
