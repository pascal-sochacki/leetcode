package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSymmetric(root *TreeNode) bool {
	return areTreesSymmetric(root.Left, root.Right)
}

func areTreesSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return right.Val == left.Val && areTreesSymmetric(left.Left, right.Right) && areTreesSymmetric(left.Right, right.Left)
}
