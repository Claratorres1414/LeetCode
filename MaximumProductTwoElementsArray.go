package main

func maxProduct(nums []int) int {
	if len(nums) == 2 {
		return (nums[0] - 1) * (nums[1] - 1)
	}

	n1 := 0
	i := 0
	n2 := 0

	for k := 0; k < len(nums); k++ {
		if nums[k] >= n1 {
			n1 = nums[k]
			i = k
		}
	}

	for k := 0; k < len(nums); k++ {
		if nums[k] >= n2 && k != i {
			n2 = nums[k]
		}
	}

	return (n1 - 1) * (n2 - 1)
}
