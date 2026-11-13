package patterns

import (
	"fmt"
	"strings"
)

// *****
// *   *
// *   *
// *   *
// *****

func P21(n int) {
	var s strings.Builder

	for i := range n {
		if i == 0 || i == n-1 {
			s.WriteString(printStar(n))
			s.WriteByte('\n')
		} else {
			for j := range n {
				if j == 0 || j == n-1 {
					s.WriteByte('*')
				} else {
					s.WriteByte(' ')
				}
			}
			s.WriteByte('\n')
		}
	}
	fmt.Print(s.String())
}

func printStar(n int) string {
	var s strings.Builder
	for range n {
		s.WriteByte('*')
	}
	return s.String()
}
