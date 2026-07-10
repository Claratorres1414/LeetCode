package main

import (
	"fmt"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	} else if divisor == 1 {
		return dividend
	} else if divisor == -1 {
		return 0 - dividend
	}
	times := 0
	if dividend > 0 && divisor > 0 {
		if dividend == divisor {
			return 1
		} else if divisor > dividend {
			return 0
		}
		for dividend >= divisor {
			dividend -= divisor
			if dividend >= divisor {
				times++
			}
		}
		return times
	} else if dividend > 0 && divisor < 0 {
		for dividend > 0 {
			dividend += divisor
			if dividend >= 0 {
				times--
			}
		}
		return times
	} else if dividend < 0 && divisor > 0 {
		for dividend < 0 {
			dividend += divisor
			if dividend <= 0 {
				times--
			}
		}
		return times
	} else {
		if dividend == divisor {
			return 1
		} else if divisor < dividend {
			return 0
		}
		for dividend <= divisor {
			dividend -= divisor
			if dividend <= divisor {
				times++
			}
		}
		return times
	}
}

func main() {
	fmt.Println(divide(-1, 1))
}
