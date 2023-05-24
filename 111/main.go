package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	right := minDepth(root.Right)
	left := minDepth(root.Left)

	if right == 0 && left == 0 {
		return 1
	} else if right == 0 {
		return left + 1
	} else if left == 0 {
		return right + 1
	} else if left > right {
		return right
	} else {
		return left
	}
}
