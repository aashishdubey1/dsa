package patterns

import (
	"fmt"
	"strconv"
	"strings"
)

// 1
// 22
// 333
// 4444
// 55555

func P4(n int) {
	var s strings.Builder
	for i := range n {
		for range i + 1 {
			s.WriteString(strconv.Itoa(i + 1))
		}
		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
