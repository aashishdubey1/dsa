package validpalindrome

import (
	"strings"
	"unicode"
)

// Given a string s, return true if it is a palindrome,
// otherwise return false.
// A palindrome is a string that reads the same forward and backward.
// It is also case-insensitive and ignores all non-alphanumeric characters.

func IsPalindrome(s string) bool {
	ms := removeNonAlpha(s)
	rs := revStr(ms)
	return rs == ms
}

func removeNonAlpha(s string) string {
	var ns strings.Builder
	for _, v := range s {
		if !unicode.IsLetter(v) {
			continue
		}
		ns.WriteRune(v)
	}
	return ns.String()
}

func revStr(str string) string {
	runes := []rune(str)
	var revs strings.Builder

	for i := len(runes) - 1; i >= 0; i-- {
		revs.WriteRune(runes[i])
	}

	return revs.String()
}
