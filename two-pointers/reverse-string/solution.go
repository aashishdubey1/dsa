package reversestring

// You are given an array of characters which represents a string s.
// Write a function which reverses a string.
// You must do this by modifying the input array in-place with O(1) extra memory.
// Input: s = ["n","e","e","t"]
// Output: ["t","e","e","n"]

func ReverseString(s []string) []string {
	pointerA := 0
	pointerB := len(s) - 1

	for pointerA < pointerB {
		s[pointerA], s[pointerB] = s[pointerB], s[pointerA]
		pointerA++
		pointerB--
	}
	return s
}
