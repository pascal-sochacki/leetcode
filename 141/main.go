package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	println(false == hasCycle(&ListNode{
		Val: 0,
	}))

	node := ListNode{
		Val: 0,
	}
	node.Next = &node
	println(true == hasCycle(&node))

	node.Next = &node
	println(false == hasCycle(&ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
		},
	}))
}

func hasCycle(head *ListNode) bool {
	var i int16 = 0
	for {
		if head.Next == nil {
			return false
		}

		if i > 10000 {
			return true
		}

		i++

		head = head.Next
	}
}
