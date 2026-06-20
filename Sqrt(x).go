package main

import "fmt"

func mySqrt(x int) int {
	i := 0
	j := 1

	for x > 0 {
		if i*i == x {
			return i
		} else if i*i < x && j*j > x {
			return i
		}

		i++
		j++
	}

	return 0
}

func main() {
	fmt.Println(mySqrt(8))
}
