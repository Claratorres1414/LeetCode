package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	r := []rune(s)
	fmt.Println(r)
	maxLen := 0
	temp := 0
	if len(r) > 1 {
		for i := 0; i < len(r); i++ {
			for j := i - 1; j >= 0; j-- {
				if r[i] == r[j] {
					if maxLen < temp {
						maxLen = temp
					}
					temp = 0
					break
				}
			}
			temp++
		}
	} else {
		maxLen = len(r)
	}
	if temp > maxLen {
		return temp
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("au"))
}
