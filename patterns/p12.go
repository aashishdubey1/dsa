package patterns

import (
	"fmt"
	"strconv"
	"strings"
)

// 1        1
// 12      21
// 123    321
// 1234  4321
// 1234554321

func P12(n int) {
	var s strings.Builder
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			s.WriteString(strconv.Itoa(j + 1))
		}

		for k := 0; k < 2*(n-i); k++ {
			s.WriteByte(' ')
		}

		for z := i; z > 0; z-- {
			s.WriteString(strconv.Itoa(z))
		}

		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
