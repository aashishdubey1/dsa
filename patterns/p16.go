package patterns

import (
	"fmt"
	"strings"
)

//
// A
// BB
// CCC
// DDDD
// EEEEE

func P16(n int) {
	var s strings.Builder
	for i := range n {
		for range i + 1 {
			s.WriteRune(rune('A' + i))
		}
		s.WriteByte('\n')
	}
	fmt.Print(s.String())
}
