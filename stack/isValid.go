package stack

func IsValid(s string) bool {
	matchingBracket := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}
	for _, v := range s {
		if expected, exist := matchingBracket[v]; exist {
			if len(stack) == 0 || expected != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}
