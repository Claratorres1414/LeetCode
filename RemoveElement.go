package main

func removeElement(nums []int, val int) int {
	k := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
			continue
		}
		k++
	}
	return k
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	println(removeElement(nums, 2))
}
