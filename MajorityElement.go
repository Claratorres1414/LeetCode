package main

import "fmt"

func majorityElement(nums []int) int {
	maxi := len(nums) / 2
	last := 0
	count := 1
	res := nums[0]

	for i := 1; count <= maxi; i++ {
		if i == len(nums) {
			i = last + 1
			res = nums[last]
			count = 1
		}

		if nums[i] == res {
			count++
		} else if last == 0 || i < last {
			last = i
		}
	}

	return res
}

func main() {
	fmt.Println(majorityElement([]int{6, 5, 5}))
}
