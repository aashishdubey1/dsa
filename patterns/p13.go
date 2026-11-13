package patterns

import (
	"fmt"
	"strconv"
	"strings"
)

// 1
// 2 3
// 4 5 6
// 7 8 9 10
// 11 12 13 14 15

func P13(n int) {
	var s strings.Builder
	var c int
	for i := range n {
		for range i + 1 {
			c++
			s.WriteString(strconv.Itoa(c))
			s.WriteByte(' ')
		}
		s.WriteByte('\n')
	}
	fmt.Print(s.String())
}
