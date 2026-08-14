package main

func maximumLengthSubstring(s string) int {
	chars := make(map[byte]int)
	var res int
	l := 0

	for r := 0; r < len(s); r++ {
		chars[s[r]]++

		for chars[s[r]] > 2 {
			chars[s[l]]--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
