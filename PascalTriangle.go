package main

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 1 {
		result[0] = []int{1}
		return result
	}
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		for j := 1; j <= i; j++ {
			if j == i {
				result[i][j] = 1
				break
			}
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
