package main

import (
	"fmt"
	"math"
)

var alfToNum = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func titleToNumber(columnTitle string) int {
	result := 0
	count := 0
	mult := 1

	for i := len(columnTitle) - 1; i >= 0; i-- {
		mult = pow(26, i)
		result += mult * alfToNum[string(columnTitle[count])]

		count++
	}

	return result
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(titleToNumber("FXSHRXW"))
}
