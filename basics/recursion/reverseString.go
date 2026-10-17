package recursion

func ReverseString(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(s[len(s)-1]) + ReverseString(s[:len(s)-1])
}
