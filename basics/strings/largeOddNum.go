package strings

import (
	"strings"
)

func LargeOddNum(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i]-'0')%2 != 0 {
			return strings.TrimLeft(s[:i+1], "0")
		}
	}
	return ""
}
