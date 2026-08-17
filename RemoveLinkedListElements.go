package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	tail := head

	for tail.Next != nil {
		if tail.Next.Val == val {
			tail.Next = tail.Next.Next
		} else {
			tail = tail.Next
		}
	}

	if head.Val == val {
		return head.Next
	}
	return head
}

func main() {
	removeElements(&ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}}, 7)
}
