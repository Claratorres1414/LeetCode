package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func inorderTraversal(root *TreeNode) []int {
	inorder(root)
	return res
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	res = append(res, root.Val)
	inorder(root.Right)
}
