package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	pass := 2
	fibonacci := 3

	for alv := 3; alv < n; alv++ {
		wait := pass
		pass = fibonacci
		fibonacci = fibonacci + wait
	}

	return fibonacci
}

func main() {
	fmt.Println(climbStairs(4))
}
