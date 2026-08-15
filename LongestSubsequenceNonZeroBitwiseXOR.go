package main

func longestSubsequence(nums []int) int {
	xorRes := xor(nums)
	res := len(nums)

	for i := 0; i < len(nums) && xorRes == 0; i++ {
		if nums[i] != 0 {
			nums[i] = 0
			res--
			xorRes = xor(nums)
		}
	}

	if xorRes != 0 {
		return res
	}

	return 0
}

func xor(nums []int) int {
	var xor int

	for _, n := range nums {
		xor ^= n
	}

	return xor
}

func main() {
	longestSubsequence([]int{0, 0, 7, 0, 0, 0, 7, 0, 0})
}
