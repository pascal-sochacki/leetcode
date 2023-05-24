package main

func main() {

	node11 := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}
	node21 := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}

	test := addTwoNumbers(&node11, &node21)

	print(test)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	val, carry := add(l1, l2, 0)
	start := ListNode{
		Val: val,
	}

	current := &start

	for {

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		val, carry = add(l1, l2, carry)

		if l1 == nil && l2 == nil && val == 0 {
			break
		}
		next := ListNode{
			Val: val,
		}
		current.Next = &next
		current = &next

	}

	return &start

}

func add(l1 *ListNode, l2 *ListNode, carry int) (int, int) {

	l1Value := 0
	if l1 != nil {
		l1Value = l1.Val
	}

	l2Value := 0
	if l2 != nil {
		l2Value = l2.Val
	}

	sum := l1Value + l2Value + carry
	return sum % 10, sum / 10
}
