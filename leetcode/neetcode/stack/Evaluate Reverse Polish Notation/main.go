package main

import (
	"fmt"
	"strconv"
)

// func evalRPN(tokens []string) int {
// 	stack := make([]int, 0)

// 	for _, token := range tokens {
// 		switch token {
// 		case "+":
// 			a := stack[len(stack)-1]
// 			b := stack[len(stack)-2]
// 			stack = stack[:len(stack)-2]
// 			stack = append(stack, b+a)
// 		case "-":
// 			a := stack[len(stack)-1]
// 			b := stack[len(stack)-2]
// 			stack = stack[:len(stack)-2]
// 			stack = append(stack, b-a)
// 		case "*":
// 			a := stack[len(stack)-1]
// 			b := stack[len(stack)-2]
// 			stack = stack[:len(stack)-2]
// 			stack = append(stack, b*a)
// 		case "/":
// 			a := stack[len(stack)-1]
// 			b := stack[len(stack)-2]
// 			stack = stack[:len(stack)-2]
// 			stack = append(stack, b/a)
// 		default:
// 			num, _ := strconv.Atoi(token)
// 			stack = append(stack, num)
// 		}
// 	}

// 	return stack[0]
// }

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		switch token {
		case "+":
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "+¥/":
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"1", "2", "+", "3", "*", "4", "-"}))

}
