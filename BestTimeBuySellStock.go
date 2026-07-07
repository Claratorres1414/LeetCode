package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	dCompra := 0
	vCompra := prices[0]
	dVenda := 1
	vVenda := prices[1]
	result := vVenda - vCompra

	for i := 0; i < len(prices); i++ {
		if prices[i] <= vCompra || prices[i] >= vVenda {
			if i < len(prices)-1 {
				if prices[i] == 0 && prices[i+1] == prices[i] {
					continue
				}
			}
			for j := i; j < len(prices); j++ {
				if prices[j]-prices[i] >= result {
					vCompra = prices[i]
					dCompra = i

					dVenda = j
					vVenda = prices[j]
					result = vVenda - vCompra
				} else if prices[j]-vCompra >= result {
					dVenda = j
					vVenda = prices[j]
					result = vVenda - vCompra
				}
			}
		}
	}

	if dCompra < dVenda && vVenda > vCompra {
		return result
	}

	return 0
}

func main() {
	fmt.Println(maxProfit([]int{0, 3, 8, 6, 8, 6, 6, 8, 2, 0, 2, 7}))
}
