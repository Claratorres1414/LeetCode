package main

var numsG []int

func sortedArrayToBST(nums []int) *TreeNode {
	numsG = nums

	return helper(0, len(nums)-1)
}

func helper(l int, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := (l + r) / 2

	root := &TreeNode{Val: numsG[mid]}
	root.Left = helper(l, mid-1)
	root.Right = helper(mid+1, r)

	return root
}
