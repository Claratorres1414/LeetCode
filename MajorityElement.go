package main

import "fmt"

func majorityElement(nums []int) int {
	count := 1
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			count++
		} else {
			count--
		}

		if count == 0 {
			res = nums[i]
			count++
		}
	}

	return res
}

func main() {
	fmt.Println(majorityElement([]int{6, 5, 5}))
}
