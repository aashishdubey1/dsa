package patterns

import (
	"fmt"
	"strings"
)

// *****
// ****
// ***
// **
// *

func P5(n int) {
	var s strings.Builder
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
