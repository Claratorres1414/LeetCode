package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i := len(nums) / 2; i > 0 && i < len(nums); {
		if nums[i] == target {
			return i
		} else if target < nums[i] {
			if i-1 >= 0 && nums[i-1] < target {
				return i
			}
			i = i / 2
		} else if target > nums[i] {
			if i+1 < len(nums) && target < nums[i+1] {
				return i + 1
			}
			i++
		}
	}

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	return 0
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}
