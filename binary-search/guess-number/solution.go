package guessnumber

import "math/rand"

var guessNumber = rand.Intn(20)


func guess(n int)int{
	if n > guessNumber {
		return 1 
	}else if n < guessNumber {
		return -1 
	}else {
		return 0 
	}
}


func GuessNumber(n int) int {
	left := 1 
	right := n 

	for left <= right {
		mid := left + (right-left)/2

		result := guess(mid)

		if result == 0 {
			return mid 
		}

		if result == -1 {
			left = mid + 1 
		}else if result == 1 {
			right = mid - 1 
		}
	}
	return -1 
}