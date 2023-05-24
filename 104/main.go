package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	right := maxDepth(root.Right)
	left := maxDepth(root.Left)

	if right > left {
		return 1 + right
	} else {
		return 1 + left
	}
}
