package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	result := mergeTwoLists(
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
	)
	println(result)

	result = mergeTwoLists(
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
	)
	println(result)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var result *ListNode
	if list1.Val > list2.Val {
		result = list2
		list2 = list2.Next
	} else {
		result = list1
		list1 = list1.Next
	}

	current := result

	for {
		if list1 == nil && list2 == nil {
			break
		} else if list1 == nil {
			current.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			current.Next = list1
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			current.Next = list2
			list2 = list2.Next
		} else {
			current.Next = list1
			list1 = list1.Next
		}
		current = current.Next
	}
	return result
}
