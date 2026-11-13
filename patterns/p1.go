package patterns

import (
	"fmt"
	"strings"
)

// Given an integer n. You need to recreate the pattern given below for any value of N.
// Let's say for N = 5, the pattern should look like as below:
//
// *****
// *****
// *****
// *****
// *****
//
// Print the pattern in the function given to you.

func P1(n int) {
	var p strings.Builder
	for range n {
		for range n {
			p.WriteString("*")
		}
		p.WriteString("\n")
	}
	fmt.Println(p.String())
}
