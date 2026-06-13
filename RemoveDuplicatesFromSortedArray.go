package main

import "fmt"

func removeDuplicates(nums []int) int {
	var res []int
	val := nums[0]
	res = append(res, val)
	count := 1
	for i := 1; i < len(nums); i++ {
		if val < nums[i] {
			count++
			res = append(res, nums[i])
		}
		val = nums[i]
	}
	for i := 0; i < len(res); i++ {
		nums[i] = res[i]
	}
	return count
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
