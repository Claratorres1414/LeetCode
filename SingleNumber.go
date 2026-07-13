package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0

	for _, n := range nums {
		res ^= n
	}

	return res
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
