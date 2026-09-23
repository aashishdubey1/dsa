package stack

import "strconv"

func CalcBaseballPoints(s []string) int {
	stack := []int{}
	for _, token := range s {

		switch token {
		case "+":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			res := a + b
			stack = append(stack, res)
		case "D":
			a := stack[len(stack)-1]
			stack = append(stack, a*2)
		case "C":
			stack = stack[:len(stack)-1]
		default:
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}

	sum := 0
	for _, val := range stack {
		sum += val
	}
	return sum
}
