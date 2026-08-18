package main

func largestInteger(nums []int, k int) int {
	appears := make(map[int]int)
	res := -1
	times := 1

	if len(nums) == k {
		res = nums[0]
		for _, n := range nums {
			res = max(res, n)
		}
		return res
	}

	for i := 0; i < len(nums); i++ {
		for j := i; j < i+k && i+k <= len(nums); j++ {
			appears[nums[j]]++
		}
	}

	for _, num := range nums {
		if num > res && appears[num] <= times {
			res = num
			times = appears[num]
		}
	}

	return res
}

func main() {
	println(largestInteger([]int{0, 0}, 2))
}
