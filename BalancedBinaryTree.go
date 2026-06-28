package main

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	resp, _ := checkBalanced(root)
	return resp
}

func checkBalanced(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftB, leftH := checkBalanced(root.Left)
	rightB, rightH := checkBalanced(root.Right)

	balance := leftB && rightB && math.Abs(float64(leftH-rightH)) <= 1

	return balance, int(1 + math.Max(float64(leftH), float64(rightH)))
}
