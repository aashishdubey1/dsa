package validpalindrome

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func ValidPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if (s[left] < 'a' || s[left] > 'z') && (s[left] < 'A' || s[left] > 'Z') {
			left++
			continue
		}
		if (s[right] < 'a' || s[right] > 'z') && (s[right] < 'A' || s[right] > 'Z') {
			right--
			continue
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}
