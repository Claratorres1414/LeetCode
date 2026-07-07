package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxP = max(maxP, profit)
		} else {
			l = r
		}
		r++
	}

	return maxP
}

func main() {
	fmt.Println(maxProfit([]int{0, 3, 8, 6, 8, 6, 6, 8, 2, 0, 2, 7}))
}
