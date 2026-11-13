package patterns

import (
	"fmt"
	"strings"
)

// A
// AB
// ABC
// ABCD
// ABCDE

func P14(n int) {
	var s strings.Builder
	for i := range n {
		for j := range i + 1 {
			s.WriteRune(rune('A' + j))
			s.WriteByte(' ')
		}
		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
