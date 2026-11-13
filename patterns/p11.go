package patterns

import (
	"fmt"
	"strings"
)

// 1
// 0 1
// 1 0 1
// 0 1 0 1
// 1 0 1 0 1

func P11(n int) {
	var s strings.Builder

	for i := range n {
		for j := range i + 1 {
			if (i+j)%2 == 0 {
				s.WriteByte('1')
			} else {
				s.WriteByte('0')
			}
		}
		s.WriteByte('\n')
	}
	fmt.Print(s.String())
}
