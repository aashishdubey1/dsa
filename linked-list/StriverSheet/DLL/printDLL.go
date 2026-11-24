package dll

import "fmt"

func PrintDLL(head *ListNode) {
	curr := head
	fmt.Print("nil")
	for curr != nil {
		fmt.Print("<-")
		fmt.Print(curr.Value)
		fmt.Print("->")
		curr = curr.Next
	}
	fmt.Println("nil")
}
