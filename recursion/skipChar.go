package recursion

import "strings"

func SkipChar(str string) string {
	var resStr strings.Builder
	return skipChar(str, &resStr)
}

func skipChar(str string, resStr *strings.Builder) string {

	if len(str) <= 0 {
		return resStr.String()
	}
	if string(str[0]) != "a" {
		resStr.WriteByte(str[0])
	}
	return skipChar(str[1:], resStr)
}
