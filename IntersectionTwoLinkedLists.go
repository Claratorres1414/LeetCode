package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tailB := headB
	var result *ListNode

	for headA != nil {
		if headA == tailB {
			result = tailB
			break
		} else if tailB == nil {
			headA = headA.Next
			tailB = headB
			continue
		}
		tailB = tailB.Next
	}

	return result
}

func main() {
	getIntersectionNode(&ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
	})
}
