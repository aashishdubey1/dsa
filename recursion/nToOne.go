package recursion

import "fmt"

func NToOne(n int) {
	if n < 1 {
		return
	}
	fmt.Println(n)
	NToOne(n - 1)
}

// no need to return here , we can also do without returning anything
// we just have to print right
