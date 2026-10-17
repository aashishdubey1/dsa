package recursion

func NumberOfStep(n int) int {

	return countSteps(n, 0)
}

func countSteps(n, step int) int {
	if n <= 0 {
		return step
	}
	if n%2 == 0 {
		return countSteps(n/2, step+1)
	}
	return countSteps(n-1, step+1)
}
