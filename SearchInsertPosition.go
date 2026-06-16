package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		}
	}

	return l
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}
