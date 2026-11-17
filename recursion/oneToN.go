package recursion

import "fmt"

func OneToN(n int) {
	if n < 1 {
		return
	}
	OneToN(n - 1)
	fmt.Println(n)
}
