package sqrt

func Sqrt(x int) int {

	left := 1 
	right := x 

	for left <= right {
		mid := left + (right-left)/2
		sqaure := mid * mid 

		if sqaure == x {
			return mid 
		}

		if sqaure > x {
			right = mid - 1 
		} else if sqaure < x {
			left = mid + 1 
		}
	}
	return left - 1 
}