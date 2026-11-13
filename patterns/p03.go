package patterns

// 1
// 12
// 123
// 1234
// 12345

import (
	"fmt"
	"strconv"
	"strings"
)

func P3(n int) {
	var s strings.Builder
	for i := range n {
		for j := range i + 1 {
			s.WriteString(strconv.Itoa(j + 1))
		}
		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
