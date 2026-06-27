package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var dMax = 0
var dist = 0

func maxDepth(root *TreeNode) int {
	dMax = 0
	dist = 0
	if root == nil {
		return 0
	}
	findMaxDepth(root)
	if dMax < dist {
		dMax = dist
	}
	return dMax
}

func findMaxDepth(node *TreeNode) {
	if node == nil {
		if dist > dMax {
			dMax = dist
		}
		dist = 0
		return
	}
	dist++
	findMaxDepth(node.Left)
	dist--
	findMaxDepth(node.Right)
	dist++
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(maxDepth(root))
}
