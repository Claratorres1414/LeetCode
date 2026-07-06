package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	triangle := build(rowIndex + 1)

	return triangle[rowIndex]
}

func build(length int) [][]int {
	result := make([][]int, length)
	for i := range result {
		result[i] = make([]int, i+1)
		for j := range result[i] {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}
	return result
}
