package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i := len(nums) / 2; i > 0 && i < len(nums); {
		if nums[i] == target {
			return i
		} else if target < nums[i] {
			i = i / 2
		} else if target > nums[i] {
			i++
		}
	}

	return 0
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}
