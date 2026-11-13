package patterns

import (
	"fmt"
	"strings"
)

// *
// **
// ***
// ****
// *****
// ****
// ***
// **
// *

func P10(n int) {
	fmt.Print(upperP(n) + downP(n-1))
}

func upperP(n int) string {
	var s strings.Builder
	for i := range n {
		for range i + 1 {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}
	return s.String()
}

func downP(n int) string {
	var s strings.Builder
	for i := n; i > 0; i-- {
		for range i {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}
	return s.String()
}
