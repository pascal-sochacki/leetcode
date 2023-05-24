package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (receiver ListNode) print() {
	temp := &receiver
	for temp != nil {
		print(temp.Val, " ")
		temp = temp.Next
	}
}

func main() {
	test := reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2}})
	print(test)
}

func reverseList(head *ListNode) *ListNode {

	var last *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = last
		last = current
		current = next
	}

	return last

}
