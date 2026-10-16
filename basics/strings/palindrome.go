package strings

func IsPalindrome(s string) bool {
	return s == revStr(s)
}

func revStr(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
