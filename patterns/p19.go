package patterns

import (
	"fmt"
	"strings"
)

// **********
// ****  ****
// ***    ***
// **      **
// *        *
// *        *
// **      **
// ***    ***
// ****  ****
// **********

func P19(n int) {
	var s strings.Builder

	// Upper half
	for i := range n {
		for j := n; j > i; j-- {
			s.WriteByte('*')
		}
		for k := 0; k < i*2; k++ {
			s.WriteByte(' ')
		}
		for j := n; j > i; j-- {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}

	// Lower half

	for i := range n {
		for range i + 1 {
			s.WriteByte('*')
		}
		for k := 0; k < 2*(n-i-1); k++ {
			s.WriteByte(' ')
		}
		for range i + 1 {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}

	fmt.Print(s.String())
}
