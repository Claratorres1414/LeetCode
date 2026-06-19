package main

import "strconv"

func addBinary(a string, b string) string {
	if len(b) == 0 {
		return a
	}

	s := ""
	carry := 0

	for c := max(len(a), len(b)) - 1; c >= 0; c-- {
		digitA := 0
		digitB := 0

		if c < len(a) {
			digitA = int(a[c])
		}

		if c < len(b) {
			digitB = int(b[c])
		}

		total := carry + digitA + digitB
		char := strconv.Itoa(total % 2)
		s = char + s
		carry = total / 2
	}

	if carry > 0 {
		s = "1" + s
	}
	return s
}
