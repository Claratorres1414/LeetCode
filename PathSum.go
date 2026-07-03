package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)
}

func dfs(node *TreeNode, curSum int, targetSum int) bool {
	if node == nil {
		return false
	}

	curSum += node.Val
	if node.Left == nil && node.Right == nil {
		return curSum == targetSum
	}

	return dfs(node.Left, curSum, targetSum) || dfs(node.Right, curSum, targetSum)
}
