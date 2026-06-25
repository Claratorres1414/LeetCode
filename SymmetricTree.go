package main

import "fmt"

func isSymmetric(root *TreeNode) bool {
	return auxOut(root.Left, root.Right) && auxIn(root.Left, root.Right)
}

func auxOut(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	} else if l.Val != r.Val {
		return false
	}

	return l.Val == r.Val && auxOut(l.Left, r.Right) && auxIn(l.Left, r.Right)
}

func auxIn(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	} else if l.Val != r.Val {
		return false
	}

	return l.Val == r.Val && auxIn(l.Right, r.Left) && auxOut(l.Right, r.Left)
}

func main() {
	root := &TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val: 25,
			Right: &TreeNode{
				Val: -95,
				Left: &TreeNode{
					Val: -100,
				},
			},
		},
		Right: &TreeNode{
			Val: 25,
			Left: &TreeNode{
				Val: -95,
				Right: &TreeNode{
					Val: -15,
				},
			},
		},
	}
	fmt.Println(isSymmetric(root))
}
