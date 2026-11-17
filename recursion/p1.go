package recursion

import "fmt"

func P1(n int) {
	printPattern(n, 0)
}

func printPattern(r, c int) {
	if r == 0 {
		return
	}

	if c < r {
		printPattern(r, c+1)
		fmt.Print("*")
	} else {
		printPattern(r-1, 0)
		fmt.Println()
	}

}
