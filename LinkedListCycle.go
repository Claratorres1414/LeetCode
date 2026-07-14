package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	s := head
	f := head

	for f != nil {
		s = s.Next
		f = f.Next
		if f == nil {
			return false
		}
		f = f.Next
		if s == f {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(hasCycle(&ListNode{
		Val: 1,
	}))
}
