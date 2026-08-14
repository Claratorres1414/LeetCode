package main

func maximumLengthSubstring(s string) int {
	chars := make(map[string]int)
	var res int
	l, r := 0, 0

	for r < len(s) {
		chars[string(s[r])]++
		r++
		l++
		res++

		if r < len(s) && chars[string(s[r])] > 2 {
			res--
			l--
		}
	}

	return res
}
