package main

import "fmt"

func singleNumber(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i != j {
				break
			} else if j == len(nums)-1 {
				return nums[i]
			}
		}
	}

	return res
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
