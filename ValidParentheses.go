package main

func isValid(s string) bool {
	parentheses := []rune(s)
	if len(parentheses)%2 != 0 {
		return false
	}
	for i := 0; i < len(parentheses); i++ {
		index := i
		char := parentheses[index]
		closure := len(parentheses) - 1
		for j := closure; j > index; j-- {
		}
	}
	return true
}
