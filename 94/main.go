package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := TreeNode{Val: 3}

	treeNode := TreeNode{Val: 2, Left: &node}

	root := TreeNode{Val: 1, Right: &treeNode}

	ints := inorderTraversal(&root)
	println(ints)

}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	right := inorderTraversal(root.Right)
	left := inorderTraversal(root.Left)

	ints := append([]int{root.Val}, right...)

	return append(ints, left...)
}
