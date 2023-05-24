package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	println(getIntersectionNode(nil, nil) == nil)
	println(getIntersectionNode(&ListNode{Val: 3}, &ListNode{Val: 3}) == nil)

	end := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 5,
		},
	}
	println(getIntersectionNode(
		&ListNode{
			Val:  3,
			Next: end,
		},
		&ListNode{
			Val:  3,
			Next: end,
		}) == end)

	end = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 5,
		},
	}

	println(getIntersectionNode(
		&ListNode{
			Val:  3,
			Next: end,
		},
		&ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  7,
				Next: end,
			},
		}) == end)

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	aTemp := headA
	var a uint16 = 0
	bTemp := headB
	var b uint16 = 0
	for {
		if aTemp.Next == nil && bTemp.Next == nil {
			if aTemp != bTemp {
				return nil
			} else {
				break
			}
		}

		if aTemp == bTemp {
			return aTemp
		}

		if aTemp.Next != nil {
			a++
			aTemp = aTemp.Next
		}

		if bTemp.Next != nil {
			b++
			bTemp = bTemp.Next
		}
	}

	aTemp = headA
	bTemp = headB

	for a > b {
		aTemp = aTemp.Next
		a--
	}

	for b > a {
		bTemp = bTemp.Next
		b--
	}

	for {
		if aTemp == bTemp {
			return aTemp
		}

		aTemp = aTemp.Next
		bTemp = bTemp.Next
	}

}
