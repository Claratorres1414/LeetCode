package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			continue
		}
		digits[i]++
		return digits
	}

	if digits[0] == 0 {
		arr := []int{1}
		for i := 0; i < len(digits); i++ {
			arr = append(arr, digits[i])
		}
		return arr
	}

	return digits
}
