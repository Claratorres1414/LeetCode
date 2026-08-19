package main

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	rows := make(map[int]int)

	for _, seat := range reservedSeats {
		row := seat[0]
		col := seat[1]

		rows[row] |= 1 << col
	}

	left := (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5)
	middle := (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7)
	right := (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9)

	res := (n - len(rows)) * 2

	for _, reserved := range rows {
		leftFree := reserved&left == 0
		middleFree := reserved&middle == 0
		rightFree := reserved&right == 0

		if leftFree && rightFree {
			res += 2
		} else if leftFree || middleFree || rightFree {
			res++
		}
	}

	return res
}
