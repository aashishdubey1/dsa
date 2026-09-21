package editor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Value string
	Prev  *Node
	Next  *Node
}

type Editor struct {
	Head *Node
	Tail *Node
	Curr *Node
}

func (editor *Editor) Type(newText string) {
	newNode := &Node{Value: newText}
	if editor.Head == nil {
		editor.Head = newNode
		editor.Tail = newNode
		editor.Curr = newNode
		return
	}

	editor.Curr.Next = newNode
	newNode.Prev = editor.Curr
	editor.Curr = newNode
	editor.Tail = newNode
}

func (editor *Editor) Undo() bool {
	if editor.Curr == nil || editor.Curr.Prev == nil {
		return false
	}
	editor.Curr = editor.Curr.Prev
	return true
}

func (editor *Editor) Redo() bool {
	if editor.Curr == nil || editor.Curr.Next == nil {
		return false
	}
	editor.Curr = editor.Curr.Next
	return true
}

func (editor *Editor) Show() {
	if editor.Curr == nil {
		fmt.Println("nothing typed yet ")
		return
	}
	fmt.Println(editor.Curr.Value)
}

func RunEditor() {
	editor := &Editor{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		switch line {
		case "u":
			ok := editor.Undo()
			if !ok {
				fmt.Println("nothing's here")
				break
			}
			editor.Show()
		case "r":
			ok := editor.Redo()
			if !ok {
				fmt.Println("nothing's here ")
				break
			}
			editor.Show()
		default:
			editor.Type(line)
			editor.Show()
		}
	}
}
