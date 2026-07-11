package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1

	for l < r {
		if !isAlfanumerico(s[l]) {
			l++
			continue
		}

		if !isAlfanumerico(s[r]) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlfanumerico(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
