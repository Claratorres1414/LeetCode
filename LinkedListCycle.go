package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	
	tail := head
	var arr []int
	for tail != nil {
		for i := 0; i < len(arr); i++ {
			if arr[i] == tail.Val {
				return true
			}
		}
		arr = append(arr, tail.Val)
		tail = tail.Next
	}

	return false
}

func main() {
	fmt.Println(hasCycle(&ListNode{
		Val: 1,
	}))
}
