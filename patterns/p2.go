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

func P2(n int) {
	var s strings.Builder
	for i := range n {
		for range i + 1 {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
