package patterns

import (
	"fmt"
	"strings"
)

// E
// D E
// C D E
// B C D E
// A B C D E

func P18(n int) {

	var s strings.Builder

	for i := range n {
		for j := i; j >= 0; j-- {
			s.WriteRune(rune('E' - j))
			s.WriteByte(' ')
		}
		s.WriteByte('\n')
	}
	fmt.Print(s.String())
}
