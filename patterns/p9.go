package patterns

import (
	"fmt"
	"strings"
)

//     *
//    ***
//   *****
//  *******
// *********
// *********
//  *******
//   *****
//    ***
//     *

func P9(n int) {
	fmt.Print(upperPyramid(n) + downPyramid(n))
}

func upperPyramid(n int) string {
	var s strings.Builder
	for i := range n {
		for j := 0; j < n-i-1; j++ {
			s.WriteByte(' ')
		}
		for j := 0; j < 2*i+1; j++ {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}
	return s.String()
}

func downPyramid(n int) string {
	var s strings.Builder
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n-i-1; j++ {
			s.WriteByte(' ')
		}
		for j := 0; j < 2*i+1; j++ {
			s.WriteByte('*')
		}
		s.WriteByte('\n')
	}
	return s.String()
}
