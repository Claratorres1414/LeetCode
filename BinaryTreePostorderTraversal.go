package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	res = visitPostorder(root, res)

	return res
}

func visitPostorder(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	res = visitPostorder(root.Left, res)
	res = visitPostorder(root.Right, res)
	res = append(res, root.Val)

	return res
}
