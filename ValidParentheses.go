package main

import "fmt"

func isValid(s string) bool {
	parenthesesMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	parentheses := []rune(s)
	var stack []rune

	for i := 0; i < len(parentheses); i++ {
		char := parentheses[i]
		closed, ok := parenthesesMap[char]
		if ok {
			if len(stack) == 0 {
				return false
			}
			stackTop := len(stack) - 1
			if stack[stackTop] == closed {
				stack = stack[:stackTop]
				continue
			}
			return false
		}
		stack = append(stack, parentheses[i])
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid("]"))
}
