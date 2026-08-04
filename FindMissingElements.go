package main

import "sort"

func findMissingElements(nums []int) []int {
	sort.Ints(nums)
	var result []int

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] {
			for j := nums[i] + 1; j < nums[i+1]; j++ {
				result = append(result, j)
			}
		}
	}

	return result
}
