package main

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	println(removeElement(nums, 2))
}
