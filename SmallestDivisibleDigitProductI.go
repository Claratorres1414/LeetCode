package main

func smallestNumber(n int, t int) int {
	if n == 10 || n == 100 {
		return n
	}

	if n > 10 {
		dec := n / 10
		uni := n % 10

		for (dec*uni)%t != 0 {
			if uni < 9 {
				uni++
			} else {
				uni = 0
				dec++
			}
		}

		return (dec * 10) + uni
	}

	for n%t != 0 && n != 10 {
		n++
	}

	return n
}
