package main

func isHappy(n int) bool {
	if n < 7 && n != 1 {
		return false
	} else if n == 1 || n == 10 || n == 100 || n == 1000 || n == 10000 || n == 100000 {
		return true
	}

	var res int

	for n > 0 {
		digit := n % 10
		res += digit * digit
		n /= 10
	}

	return isHappy(res)
}
