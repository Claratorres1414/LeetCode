package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return minimum(root) + 1
}

func minimum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return min(minDepth(root.Left), minDepth(root.Right))
}
