package main

import (
	"fmt"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	result := &ListNode{Val: head.Val}
	tail := result
	head = head.Next

	for head != nil {
		if head.Val == tail.Val {
			head = head.Next
			continue
		}

		tail.Next = &ListNode{Val: head.Val}
		head = head.Next
		tail = tail.Next
	}

	return result
}

func main() {
	fmt.Println(deleteDuplicates(
		&ListNode{0,
			&ListNode{0,
				&ListNode{0,
					&ListNode{0, nil},
				},
			},
		},
	),
	)
}
