package strings

import "fmt"

// Given a string, the task is to reverse it. The string is represented by an array of characters s.
// Perform the reversal in place with O(1) extra memory.
// Note: no need to return anything, modify the given list.

func ReverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	fmt.Println(string(s))
}
