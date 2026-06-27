package main

var res []int

func inorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
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
