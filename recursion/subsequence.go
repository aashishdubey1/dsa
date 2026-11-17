package recursion

import (
	"fmt"
)

func Subsequence(str string) {
	var p string
	printSubsequence(p, str)
}

func printSubsequence(p, up string) {
	if len(up) <= 0 {
		fmt.Println(p)
		return
	}

	printSubsequence(p+string(up[0]), up[1:])
	printSubsequence(p, up[1:])
}
