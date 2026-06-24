package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return verify(p, q)
}

func verify(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	return verify(p.Left, q.Left) && verify(p.Right, q.Right)
}
