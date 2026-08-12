package main

func maxSubarrayLength(nums []int, k int) int {
	freq := make(map[int]int)

	left := 0
	maxLength := 0

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++

		for freq[nums[right]] > k {
			freq[nums[left]]--
			left++
		}

		length := right - left + 1

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
