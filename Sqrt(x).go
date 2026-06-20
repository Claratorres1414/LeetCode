package main

import "fmt"

func mySqrt(x int) int {
	l := 0
	r := x
	result := 0

	for l <= r {
		mid := l + ((r - l) / 2)

		if mid*mid > x {
			r = mid - 1
		} else if mid*mid < x {
			l = mid + 1
			result = mid
		} else {
			return mid
		}
	}

	return result
}

func main() {
	fmt.Println(mySqrt(8))
}
