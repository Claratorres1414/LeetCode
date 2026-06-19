package main

func addBinary(a string, b string) string {
	if len(b) == 0 {
		return a
	}

	aRunes := []rune(a)
	bRunes := []rune(b)
	s := ""
	carry := 0

	if len(aRunes) > len(bRunes) {
		pointer := len(bRunes) - 1
		for i := len(aRunes) - 1; i >= 0; i-- {
			if aRunes[i] != bRunes[pointer] && carry == 0 {
				aRunes[i] = '1'
			} else if aRunes[i] == '1' && carry == 0 {
				aRunes[i] = '0'
				carry = 1
			}
		}
	}

	return s
}
