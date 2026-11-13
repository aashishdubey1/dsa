package patterns

import (
	"fmt"
	"strings"
)

// ABCDE
// ABCD
// ABC
// AB
// A

func P15(n int) {
	var s strings.Builder
	for i := range n {
		for j := 0; j < n-i; j++ {
			s.WriteRune(rune('A' + j))
			s.WriteByte(' ')
		}
		s.WriteByte('\n')
	}
	fmt.Print(s.String())
}
