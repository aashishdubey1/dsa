package stack

import "strconv"

func EvalRPN(s []string) int {

	stack := ArrayStack{}
	for _, token := range s {
		switch token {
		case "+":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			res := a + b
			stack.Push(res)
		case "-":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			res := a - b
			stack.Push(res)
		case "*":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			res := a * b
			stack.Push(res)
		case "/":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			res := a / b
			stack.Push(res)
		default:
			res, _ := strconv.Atoi(token)
			stack.Push(res)
		}
	}
	res, _ := stack.Pop()
	return res
}
