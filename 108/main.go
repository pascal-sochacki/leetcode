package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	bst := sortedArrayToBST([]int{})
	bst = sortedArrayToBST([]int{1, 2, 3})

	println(bst)
}

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	var half int = length / 2

	return &TreeNode{
		Val:   nums[half],
		Left:  sortedArrayToBST(nums[:half]),
		Right: sortedArrayToBST(nums[half+1:]),
	}
}
