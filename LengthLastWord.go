package main

func lengthOfLastWord(s string) int {
	val := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			val++
		} else if s[i] == ' ' && val != 0 {
			return val
		} else if s[i] == ' ' && val == 0 {
			continue
		}
	}
	return val
}
