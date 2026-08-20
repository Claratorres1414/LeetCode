package main

func resultArray(nums []int) []int {
	var arr1 []int
	var arr2 []int

	arr1 = append(arr1, nums[0])
	arr2 = append(arr2, nums[1])

	i1 := 0
	i2 := 0

	for i := 2; i < len(nums); i++ {
		if arr1[i1] > arr2[i2] {
			arr1 = append(arr1, nums[i])
			i1++
		} else {
			arr2 = append(arr2, nums[i])
			i2++
		}
	}

	return append(arr1, arr2...)
}
