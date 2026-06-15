package main

import "fmt"

func strStr(haystack string, needle string) int {
	count := 0
	index := 0
	nr := []rune(needle)
	hr := []rune(haystack)

	for i := 0; i < len(hr); i++ {
		if hr[i] != nr[count] {
			count = 0
			i = index
			index++
			continue
		}
		count++
		if count == 1 {
			index = i
		}
		if count == len(nr) {
			return index
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}
