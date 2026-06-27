package main

func sortedArrayToBST(nums []int) *TreeNode {
	mid := len(nums) / 2
	l := mid - 1
	r := mid + 1
	bst := &TreeNode{Val: nums[mid]}
	lNode := bst
	rNode := bst
	for l >= 0 || r < len(nums) {
		if l >= 0 {
			lNode.Left = &TreeNode{Val: nums[l]}
			lNode = lNode.Left
			l--
		}
		if r < len(nums) {
			rNode.Right = &TreeNode{Val: nums[r]}
			rNode = rNode.Right
			r++
		}
	}
	return bst
}
