package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1

	carry := 0

	res := ""

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		res = string(byte(sum%2)+'0') + res

		carry = sum / 2
	}

	return res
}

func main() {
	fmt.Println(addBinary("1010", "1011"))
}
