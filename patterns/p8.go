package patterns

import (
	"fmt"
	"strings"
)

// *********
//  *******
//   *****
//    ***
//     *

func P8(n int) {
	var s strings.Builder
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n-i-1; j++ {
			s.WriteByte(' ')
		}

		for j := 0; j < 2*i+1; j++ {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
