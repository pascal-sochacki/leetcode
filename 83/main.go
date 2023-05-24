package main

func main() {
	node := deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}})
	println(node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	temp := head
	lastChange := head

	for temp != nil {

		if lastChange.Val != temp.Val {
			lastChange.Next = temp
			lastChange = temp
		}

		temp = temp.Next
	}

	if lastChange.Next != nil {
		lastChange.Next = nil
	}

	return head
}
