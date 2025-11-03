package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	stack := make([]rune, 0, len(s))

	ope := 0
	for _, d := range s {

		switch d {
		case '(':
			ope++
			stack = append(stack, d)
		case '[':
			ope++
			stack = append(stack, d)
		case '{':
			ope++
			stack = append(stack, d)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				ope--
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				ope--
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				ope--
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return ope == 0 && len(stack) == 0
}

func main() {
	fmt.Println(isValid("([{}])"))
	fmt.Println(isValid("[(])"))
	fmt.Println(isValid("]]"))
}

// シンプルなたりかた
// mapで open closeを持っていて pop した値でが mapで該当すればOKみたいな感じ
// func isValid(s string) bool {
//     stack := linkedliststack.New()
//     closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

//     for _, c := range s {
//         if open, exists := closeToOpen[c]; exists {
//             if !stack.Empty() {
//                 top, ok := stack.Pop()
//                 if ok && top.(rune) != open {
//                     return false
//                 }
//             } else {
//                 return false
//             }
//         } else {
//             stack.Push(c)
//         }
//     }

//     return stack.Empty()
// }
