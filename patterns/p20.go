package patterns

import (
	"fmt"
	"strings"
)

// *        *
// **      **
// ***    ***
// ****  ****
// **********
// ****  ****
// ***    ***
// **      **
// *        *

func P20(n int) {
	var s strings.Builder

	// Upper half
	for i := range n {
		for range i + 1 {
			s.WriteByte('*')
		}
		for j := 0; j < 2*(n-i-1); j++ {
			s.WriteByte(' ')
		}
		for range i + 1 {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}
	// lower Half
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			s.WriteByte('*')
		}
		for k := 0; k < 2*(n-i); k++ {
			s.WriteByte(' ')
		}
		for j := 0; j < i; j++ {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
