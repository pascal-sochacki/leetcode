package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	println(isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	heightDiff := height(root.Right) - height(root.Left)
	if heightDiff > 1 || heightDiff < -1 {
		return false
	}

	return isBalanced(root.Right) && isBalanced(root.Left)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	right := height(root.Right)
	left := height(root.Left)

	if right > left {
		return right + 1
	} else {
		return left + 1
	}

}
