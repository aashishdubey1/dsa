package patterns

import (
	"fmt"
	"strconv"
	"strings"
)

// 12345
// 1234
// 123
// 12
// 1

func P6(n int) {
	var s strings.Builder
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			s.WriteString(strconv.Itoa(j + 1))
		}
		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
