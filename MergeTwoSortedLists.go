package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	tail := result

	for list1 != nil || list2 != nil {
		if list1 == nil {
			tail.Next = list2
			return result.Next
		}
		if list2 == nil {
			tail.Next = list1
			return result.Next
		}
		if list1.Val <= list2.Val {
			tail.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			tail.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		tail = tail.Next
	}
	return result.Next
}
